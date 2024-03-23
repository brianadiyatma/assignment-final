package service

import (
	dto "assignment-final/model/DTO"
	"assignment-final/model/domain"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PhotoService struct {
	DB *gorm.DB
}

func NewPhotoService(db *gorm.DB) *PhotoService {
	return &PhotoService{
		DB: db,
	}
}

func (ps *PhotoService) CreatePhoto(createPhotoDTO dto.CreatePhotoDTO, userID uuid.UUID) (dto.PhotoResponseDTO, error) {
	photo := domain.Photo{
		Title:    createPhotoDTO.Title,
		Caption:  createPhotoDTO.Caption,
		PhotoUrl: createPhotoDTO.PhotoUrl,
		UserID:   userID,
	}

	result := ps.DB.Create(&photo)
	if result.Error != nil {
		return dto.PhotoResponseDTO{}, result.Error
	}

	response := dto.PhotoResponseDTO{
		ID:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoUrl:  photo.PhotoUrl,
		UserID:    photo.UserID,
		CreatedAt: photo.CreatedAt,
	}

	return response, nil
}

func (ps *PhotoService) GetAllPhotos() ([]domain.Photo, error) {
	var photos []domain.Photo
	if result := ps.DB.Preload("User").Find(&photos); result.Error != nil {
		return nil, result.Error
	}

	return photos, nil
}



func (s *PhotoService) EditPhoto(photoID uuid.UUID, editPhotoDTO dto.EditPhotoDTO) (dto.GetPhotoDTO, error) {

    tx := s.DB.Begin()
    var photo domain.Photo
    if err := tx.Preload("User").First(&photo, "id = ?", photoID).Error; err != nil {
        tx.Rollback()
        return dto.GetPhotoDTO{}, err
    }
    photo.Title = editPhotoDTO.Title
    photo.Caption = editPhotoDTO.Caption
    photo.PhotoUrl = editPhotoDTO.PhotoUrl
    if err := tx.Model(&photo).Select("Title", "Caption", "PhotoUrl").Updates(domain.Photo{Title: photo.Title, Caption: photo.Caption, PhotoUrl: photo.PhotoUrl}).Error; err != nil {

        tx.Rollback()
        return dto.GetPhotoDTO{}, err
    }
    if err := tx.Commit().Error; err != nil {
        return dto.GetPhotoDTO{}, err
    }
    responseDTO := dto.GetPhotoDTO{
        ID:        photo.ID,
        Title:     photo.Title,
        Caption:   photo.Caption,
        PhotoUrl:  photo.PhotoUrl,
        UserID:    photo.UserID,
        User: dto.UserMinimalInfoDTO{
            Username: photo.User.Username,
            Email:    photo.User.Email,
        },
        CreatedAt: photo.CreatedAt,
        UpdatedAt: photo.UpdatedAt,
    }

    return responseDTO, nil
}

func (s *PhotoService) DeletePhoto(photoID uuid.UUID, userID uuid.UUID) error {
    var photo domain.Photo

    if err := s.DB.Preload("User").First(&photo, "id = ?", photoID).Error; err != nil {
        return err
    }

    // Check if the photo belongs to the user making the request
    if photo.UserID != userID {
        return errors.New("unauthorized: you do not own this photo")
    }

    if err := s.DB.Delete(&photo).Error; err != nil {
        return err
    }

    return nil
}

