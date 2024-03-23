package controller

import (
	dto "assignment-final/model/DTO"
	"assignment-final/service"
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController struct {
	UserService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	var updateDTO dto.UpdateUserDTO
	if err := c.ShouldBindJSON(&updateDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uuid := c.Param("uuid")
	val, exist := c.Get("uuid")

	
	if !exist{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Terjadi Kesalahan Ketika Parsing Token"})
		return
	}
	if val != uuid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Anda Tidak Berhak Mengedit User Milik Orang Lain"})
		return
	}
	updatedUser, err := uc.UserService.UpdateUser(uuid, updateDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}


func (uc *UserController) DeleteUserHandler(c *gin.Context) {
    uuidInterface, exists := c.Get("uuid")
	uuidParams := c.Param("uuid")

    if !exists {
        c.JSON(http.StatusBadRequest, gin.H{"error": "UUID dari Token Tidak ditemukan"})
        return
    }
	

    userUUID, ok := uuidInterface.(string)
	fmt.Print(userUUID)
    if !ok {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Format UUID Salah"})
        return
    }

	if uuidParams !=userUUID {
		  c.JSON(http.StatusBadRequest, gin.H{"error": "Anda Tidak Berhak Delete User Orang"})
        return
	}

    userUUIDParsed, err := uuid.Parse(userUUID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Format UUID Tidak Benar"})
        return
    }

    err = uc.UserService.DeleteUser(userUUIDParsed)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Terjadi Kesalahan tidak dapat menghapus user"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User anda telah dihapus"})
}

