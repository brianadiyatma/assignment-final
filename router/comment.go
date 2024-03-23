package router

import (
	"assignment-final/controller"
	"assignment-final/middleware"

	"github.com/gin-gonic/gin"
)

type commentRouter struct {
	commentController controller.CommentController
}


func NewCommentRouter(cc controller.CommentController) *commentRouter {
	return &commentRouter{
		commentController: cc,
	}
}


func (cr *commentRouter) CommentRouter(r *gin.Engine) {
	comments := r.Group("/comments")
	comments.Use(middleware.AuthMiddleware)
	{
		comments.POST("/", cr.commentController.CreateCommentHandler)

		comments.GET("/", cr.commentController.GetCommentsHandler)

		comments.DELETE("/:uuid", cr.commentController.DeleteCommentHandler)
	}
}
