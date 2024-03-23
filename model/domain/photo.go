package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

	type Photo struct {
		ID        uuid.UUID   `gorm:"primaryKey"`
		Title     string `gorm:"not null"`
		Caption   string
		PhotoUrl  string `gorm:"not null"`
		UserID    uuid.UUID   `gorm:"not null"`
		User      User   `gorm:"foreignKey:UserID"`
		DeletedAt gorm.DeletedAt
		UpdatedAt time.Time
		CreatedAt time.Time
	}

	func (photo *Photo) BeforeCreate(db *gorm.DB) error {

	photo.ID = uuid.New()

	return nil
}
