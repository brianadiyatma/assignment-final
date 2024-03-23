package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateSocialMediaDTO struct {
	Name           string `json:"name" binding:"required"`
	SocialMediaUrl string `json:"social_media_url" binding:"required"`
}

type CreateSocialMediaResponseDTO struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"socialMediaUrl"`
	UserID         uuid.UUID `json:"userId"`
	UpdatedAt      time.Time `json:"updatedAt"`
	CreatedAt      time.Time `json:"createdAt"`
}

type EditSocialMediaDTO struct {
	Name           string `json:"name" binding:"required"`
	SocialMediaUrl string `json:"social_media_url" binding:"required,url"`
}
type EditSocialMediaResponseDTO struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserID         uuid.UUID `json:"user_id"`
	UpdatedAt      time.Time `json:"updated_at"`
}