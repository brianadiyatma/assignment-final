package controller

import (
	dto "assignment-final/model/DTO"
	"assignment-final/model/domain"
	"assignment-final/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PhotoController struct {
    photoService *service.PhotoService
}

func NewPhotoController(ps *service.PhotoService) *PhotoController {
	return &PhotoController{
		photoService: ps,
	}
}

func (pc *PhotoController) CreatePhoto(c *gin.Context) {
    var createPhotoDTO dto.CreatePhotoDTO
    if err := c.ShouldBindJSON(&createPhotoDTO); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userID, exists := c.Get("uuid")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "user ID not found"})
        return
    }
	parsedid , err:=  uuid.Parse(userID.(string))
	
 	if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Format UUID Salah"})
        return
    }
    photo, err := pc.photoService.CreatePhoto(createPhotoDTO,parsedid )
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, photo)
}

func (pc *PhotoController) GetPhotosHandler(c *gin.Context) {
	photos, err := pc.photoService.GetAllPhotos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Convert to DTO
	var photosResponse []dto.GetPhotoDTO
	for _, photo := range photos {
		photoResponse := dto.GetPhotoDTO{
			ID:        photo.ID,
			Title:     photo.Title,
			Caption:   photo.Caption,
			PhotoUrl:  photo.PhotoUrl,
			UserID:    photo.UserID,
			User: dto.UserMinimalInfoDTO{
				Username: photo.User.Username,
				Email:    photo.User.Email,
			},
			CreatedAt: photo.CreatedAt,
			UpdatedAt: photo.UpdatedAt,
		}
		photosResponse = append(photosResponse, photoResponse)
	}

	c.JSON(http.StatusOK, photosResponse)
}

func (pc *PhotoController) EditPhotoHandler(c *gin.Context) {
    var editDTO dto.EditPhotoDTO
    if err := c.ShouldBindJSON(&editDTO); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    photoID, _ := uuid.Parse(c.Param("uuid"))
    userIDunParsed, exists := c.Get("uuid") 
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Auth Bermasalah"})
        return
    }
	parsedUUID ,err := uuid.Parse(userIDunParsed.(string))
	if err != nil {
		        c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
	}

    var photo domain.Photo
    if err := pc.photoService.DB.First(&photo, "id = ?", photoID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Foto Tidak Ditemukan"})
        return
    }

    if photo.UserID != parsedUUID {
        c.JSON(http.StatusForbidden, gin.H{"error": "Anda Bukan Pemilik Foto Ini"})
        return
    }

    result ,err := pc.photoService.EditPhoto(photoID, editDTO)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message":"Photo telah diedit", "data":result})
}


func (pc *PhotoController) DeletePhotoHandler(c *gin.Context) {
    userUUID, exists := c.Get("uuid")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Tidak authorize"})
        return
    }

    photoUUID, err := uuid.Parse(c.Param("uuid"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "UUID Tidak Valid"})
        return
    }
    parseduserUUID, err :=uuid.Parse(userUUID.(string))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Formatnya Bermasalah"})
    }

    err = pc.photoService.DeletePhoto(photoUUID,parseduserUUID )
    if err != nil {
        if err.Error() == "Anda Bukan Pemilik Foto Ini" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Tidak Bisa Menghapus Foto"})
        }
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Photo Anda Berhasl Dihapus"})
}