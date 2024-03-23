package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateCommentDTO struct {
	Message string    `json:"message" binding:"required"`
	PhotoID uuid.UUID `json:"photo_id" binding:"required"`
}

type CreateCommentResponseDTO struct {
	ID        uuid.UUID `json:"id"`
	Message   string    `json:"message"`
	PhotoID   uuid.UUID `json:"photo_id"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type CommentResponseDTO struct {
	ID        uuid.UUID   `json:"id"`
	Message   string      `json:"message"`
	PhotoID   uuid.UUID   `json:"photo_id"`
	UserID    uuid.UUID   `json:"user_id"`
	UpdatedAt time.Time   `json:"updated_at"`
	CreatedAt time.Time   `json:"created_at"`
	User      UserDTO     `json:"user"`
}

type UserDTO struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Username string    `json:"username"`
}