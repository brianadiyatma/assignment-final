package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	ID        uuid.UUID   `gorm:"primaryKey"`
	UserID    uuid.UUID   `gorm:"not null"`
	PhotoID   uuid.UUID   `gorm:"not null"`
	Message   string `gorm:"not null"`
	User      User  `gorm:"foreignKey:UserID"`
	Photo     Photo `gorm:"foreignKey:PhotoID"`
	UpdatedAt time.Time
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt

}

	func (comment *Comment) BeforeCreate(db *gorm.DB) error {

	comment.ID = uuid.New()

	return nil
}