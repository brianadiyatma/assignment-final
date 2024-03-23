package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateUserDTO struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Age      int    `json:"age" binding:"required,gte=0"`
}

type LoginDTO struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

type GetUserDTO struct {
    ID        uuid.UUID      `json:"id"`
    Username  string    `json:"username"`
    Email     string    `json:"email"`
    Age       int       `json:"age"`
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
}

type LoginResponseDTO struct {
    Token     string    `json:"token"`
    UserID    uuid.UUID      `json:"user_id"`
    Username  string    `json:"username"`
    Email     string    `json:"email"`
    Age       int       `json:"age"`
    CreatedAt time.Time `json:"created_at"`
}
type UpdateUserDTO struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type GetAllSocialMediaDTO struct {
	ID             uuid.UUID          `json:"id"`
	Name           string             `json:"name"`
	SocialMediaUrl string             `json:"socialMediaUrl"`
	UserID         uuid.UUID          `json:"userId"`
	User           UserMinimalInfoDTO `json:"user"`
	UpdatedAt      time.Time          `json:"updatedAt"`
	CreatedAt      time.Time          `json:"createdAt"`
}