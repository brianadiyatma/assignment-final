package controller

import (
	dto "assignment-final/model/DTO"
	"assignment-final/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CommentController struct {
	commentService *service.CommentService
}

func NewCommentService (cs *service.CommentService)*CommentController {
	return &CommentController{
		commentService: cs,
	}
}

func (c *CommentController) CreateCommentHandler(ctx *gin.Context) {
    var createDTO dto.CreateCommentDTO
    if err := ctx.ShouldBindJSON(&createDTO); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    unParsedUserID, exists := ctx.Get("uuid")
    if !exists {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
        return
    }
	parsedUUID, err := uuid.Parse(unParsedUserID.(string))
	if err != nil {
		        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Parsing UUID Bermasalah"})
        return
	}

    response, err := c.commentService.CreateComment(parsedUUID, createDTO)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, response)
}

func NewCommentController(cs *service.CommentService) *CommentController {
	return &CommentController{commentService: cs}
}

func (cc *CommentController) GetCommentsHandler(c *gin.Context) {
	comments, err := cc.commentService.GetComments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comments)
}

func (c *CommentController) DeleteCommentHandler(ctx *gin.Context) {
    commentID := ctx.Param("uuid")
    unParsedUserID, exists := ctx.Get("uuid") 
    if !exists {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
        return
    }
	userID, err := uuid.Parse(unParsedUserID.(string))
	if err != nil {
		 ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Parsing UUID GAGAL"})
        return
	}
    c.commentService.DeleteComment(uuid.MustParse(commentID), userID)
    if err != nil {
        if err.Error() == "unauthorized to delete this comment" {
            ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        } else {
            ctx.JSON(http.StatusNotFound, gin.H{"error": "comment not found or error deleting the comment"})
        }
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "comment deleted successfully"})
}