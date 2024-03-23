package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SocialMedia struct {
	ID             uuid.UUID   `gorm:"primaryKey"`
	Name           string `gorm:"not null"`
	SocialMediaUrl string `gorm:"not null;type:text"`
	UserID         uuid.UUID   `gorm:"not null"`
	User           User `gorm:"foreignKey:UserID"`
	UpdatedAt      time.Time
	CreatedAt      time.Time
	DeletedAt 	   gorm.DeletedAt
}

func (photo *SocialMedia) BeforeCreate(db *gorm.DB) error {

	photo.ID = uuid.New()

	return nil
}
