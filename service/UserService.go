package service

import (
	dto "assignment-final/model/DTO"
	"assignment-final/model/domain"
	"assignment-final/util"
	"errors"
	"time"

	"github.com/google/uuid"
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

func (s *UserService) CreateUser(request dto.CreateUserDTO) (dto.GetUserDTO, error) {

	user := domain.User{
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
		Age:      request.Age,
	}

	result := s.DB.Create(&user)
	
	getUserDTO := dto.GetUserDTO{
        ID:        user.ID,
        Username:  user.Username,
        Email:     user.Email,
        Age:       user.Age,
        CreatedAt: user.CreatedAt,
    }

	return getUserDTO, result.Error
}


func (s *UserService) Login(request dto.LoginDTO) (dto.LoginResponseDTO, error) {
    var user domain.User
    result := s.DB.Where("email = ?", request.Email).First(&user)

    if result.Error != nil {
        return dto.LoginResponseDTO{}, errors.New("user not found")
    }

    if !util.CheckPasswordHash(request.Password, user.Password) {
        return dto.LoginResponseDTO{}, errors.New("incorrect password")
    }

    token, err := util.GenerateJWT(user.ID.String())
    if err != nil {
        return dto.LoginResponseDTO{}, err
    }


    loginResponseDTO := dto.LoginResponseDTO{
        Token:     token,
        UserID:    user.ID,
        Username:  user.Username,
        Email:     user.Email,
        Age:       user.Age,
        CreatedAt: user.CreatedAt,
    }

    return loginResponseDTO, nil
}



func (s *UserService) UpdateUser(uuidString string, updateDTO dto.UpdateUserDTO) (*dto.GetUserDTO, error) {
	var user domain.User
	uuid, err := uuid.Parse(uuidString)
	if err != nil {
		return nil, err
	}

	result := s.DB.Where("id = ?", uuid).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	user.Username = updateDTO.Username
	user.Email = updateDTO.Email
	user.UpdatedAt = time.Now()

	result = s.DB.Save(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &dto.GetUserDTO{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Age:       user.Age,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (s *UserService) DeleteUser(id uuid.UUID) error {
    result := s.DB.Where("id = ?", id).Delete(&domain.User{})
    return result.Error
}

func (cs *CommentService) GetComments() ([]dto.CommentResponseDTO, error) {
	var comments []domain.Comment
	err := cs.DB.Preload("User").Find(&comments).Error
	if err != nil {
		return nil, err
	}

	var commentDTOs []dto.CommentResponseDTO
	for _, comment := range comments {
		commentDTO := dto.CommentResponseDTO{
			ID:        comment.ID,
			Message:   comment.Message,
			PhotoID:   comment.PhotoID,
			UserID:    comment.UserID,
			UpdatedAt: comment.UpdatedAt,
			CreatedAt: comment.CreatedAt,
			User: dto.UserDTO{
				ID:       comment.User.ID,
				Email:    comment.User.Email,
				Username: comment.User.Username,
			},
		}
		commentDTOs = append(commentDTOs, commentDTO)
	}

	return commentDTOs, nil
}