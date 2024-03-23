package domain

import (
	"assignment-final/util"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {

	ID        uuid.UUID `gorm:"not null;uniqueIndex"`
	Username  string `gorm:"not null;uniqueIndex"`
	Email     string `gorm:"not null;uniqueIndex"`
	Password  string `gorm:"not null"`
	Age       int    `gorm:"not null"`
	UpdatedAt time.Time
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
	
}

func (user *User) BeforeCreate(db *gorm.DB) error {
	user.Password = util.Hash(user.Password)
	user.ID = uuid.New()

	return nil
}
