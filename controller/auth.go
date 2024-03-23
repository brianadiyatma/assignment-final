package controller

import (
	dto "assignment-final/model/DTO"
	"assignment-final/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	userService *service.UserService
}

func NewAuthController(userService *service.UserService) *AuthController {
	return &AuthController{
		userService: userService,
	}
}

func (ac *AuthController) CreateUserHandler(c *gin.Context) {
	var dto dto.CreateUserDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := ac.userService.CreateUser(dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": user})
}

func (ac *AuthController) LoginController(c *gin.Context) {
    var loginDTO dto.LoginDTO
    if err := c.ShouldBindJSON(&loginDTO); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    result, err := ac.userService.Login(loginDTO)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "login failed"})
        return
    }

    c.JSON(http.StatusOK, result)
}

