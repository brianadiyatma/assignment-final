package service

import (
	dto "assignment-final/model/DTO"
	"assignment-final/model/domain"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		DB: db,
	}
}

func (s *UserService) CreateUser(dto dto.CreateUserDTO) ( domain.User, error) {

	user := domain.User{
		Username: dto.Username,
		Email:    dto.Email,
		Password: dto.Password,
		Age:      dto.Age,
	}

	result := s.DB.Create(&user)
	return user, result.Error
}
