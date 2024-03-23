package controller

import (
	dto "assignment-final/model/DTO"
	"assignment-final/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SocialMediaController struct {
	Service *service.SocialMediaService
}

func NewSocialMediaController(service *service.SocialMediaService) *SocialMediaController {
	return &SocialMediaController{
		Service: service,
	}
}

func (c *SocialMediaController) CreateSocialMediaHandler(ctx *gin.Context) {
	var createDTO dto.CreateSocialMediaDTO
	if err := ctx.ShouldBindJSON(&createDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	unParsedUserID, exists := ctx.Get("uuid")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Error Auth"})
		return
	}
	parsedUUID, err := uuid.Parse(unParsedUserID.(string))
if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Parsign UUID error"})
		return
	}
	socialMedia, err := c.Service.CreateSocialMedia(parsedUUID, createDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	 responseDTO := dto.CreateSocialMediaResponseDTO{
        ID:             socialMedia.ID,
        Name:           socialMedia.Name,
        SocialMediaUrl: socialMedia.SocialMediaUrl,
        UserID:         socialMedia.UserID,
        UpdatedAt:      socialMedia.UpdatedAt,
        CreatedAt:      socialMedia.CreatedAt,
    }

	ctx.JSON(http.StatusCreated, responseDTO)
}

func (smc *SocialMediaController) GetSocialMediasHandler(c *gin.Context) {
	socialMedias, err := smc.Service.GetAllSocialMedia()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, socialMedias)
}

func (smc *SocialMediaController) EditSocialMediaHandler(c *gin.Context) {
	var editDTO dto.EditSocialMediaDTO
	if err := c.ShouldBindJSON(&editDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	unParseduserID, _ := c.Get("uuid") // Assuming "userID" is set by your authentication middleware
	socialMediaID, err := uuid.Parse(c.Param("uuid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid social media ID"})
		return
	}
	userID, err:= uuid.Parse(unParseduserID.(string))
	if err != nil {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "UUID Bermasalah"})
		return
	}

	updatedSocialMedia, err := smc.Service.EditSocialMedia(userID, socialMediaID, editDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.GetAllSocialMediaDTO{
		ID: updatedSocialMedia.ID,
		Name: updatedSocialMedia.Name,
		SocialMediaUrl: updatedSocialMedia.SocialMediaUrl,
		UserID: updatedSocialMedia.UserID,
		User: dto.UserMinimalInfoDTO{
			Username: updatedSocialMedia.User.Username,
			Email: updatedSocialMedia.User.Email,
		},
		UpdatedAt: updatedSocialMedia.UpdatedAt,
		CreatedAt: updatedSocialMedia.CreatedAt,
	})
}

func (ctrl *SocialMediaController) DeleteSocialMediaHandler(c *gin.Context) {

	userID, exists := c.Get("uuid")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Wajib Login Dulu Gan"})
		return
	}

	userUUID, err := uuid.Parse(userID.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID Tidak Valid"})
		return
	}

	socialMediaIDParam := c.Param("uuid")
	socialMediaUUID, err := uuid.Parse(socialMediaIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "UUID Tidak Valid (Sepertinya)"})
		return
	}

	err = ctrl.Service.DeleteSocialMedia(userUUID, socialMediaUUID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Item yang ingin anda delete tidak ditemukan"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus Entry Sosial Media"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Nerhasil Menghapus Social Media"})
}