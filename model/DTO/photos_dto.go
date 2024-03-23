package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreatePhotoDTO struct {
	Title    string `json:"title" binding:"required"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url" binding:"required"`
}

type PhotoResponseDTO struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption,omitempty"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"createdAt"`
}

type GetPhotoDTO struct {
	ID        uuid.UUID          `json:"id"`
	Title     string             `json:"title"`
	Caption   string             `json:"caption,omitempty"`
	PhotoUrl  string             `json:"photoUrl"`
	UserID    uuid.UUID          `json:"userId"`
	User      UserMinimalInfoDTO `json:"user"`
	CreatedAt time.Time          `json:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt"`
}

type UserMinimalInfoDTO struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type EditPhotoDTO struct {
    Title     string `json:"title"`
    Caption   string `json:"caption,omitempty"`
    PhotoUrl  string `json:"photo_url"`
}