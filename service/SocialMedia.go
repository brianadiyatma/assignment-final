package service

import (
	dto "assignment-final/model/DTO"
	"assignment-final/model/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SocialMediaService struct {
	DB *gorm.DB
}

func NewSocialMedia(db *gorm.DB) *SocialMediaService {
	return &SocialMediaService{
		DB: db,
	}
}

func (s *SocialMediaService) CreateSocialMedia(userID uuid.UUID, dto dto.CreateSocialMediaDTO) (domain.SocialMedia, error) {
	socialMedia := domain.SocialMedia{
		Name:           dto.Name,
		SocialMediaUrl: dto.SocialMediaUrl,
		UserID:         userID,
	}

	err := s.DB.Create(&socialMedia).Error
	if err != nil {
		return domain.SocialMedia{}, err
	}

	return socialMedia, nil
}

func (s *SocialMediaService) GetAllSocialMedia() ([]dto.GetAllSocialMediaDTO, error) {
    var socialMedias []domain.SocialMedia
    var responseDTOs []dto.GetAllSocialMediaDTO
    result := s.DB.Preload("User").Find(&socialMedias)
    if result.Error != nil {
        return nil, result.Error
    }


    for _, socialMedia := range socialMedias {
        responseDTO := dto.GetAllSocialMediaDTO{
            ID:             socialMedia.ID,
            Name:           socialMedia.Name,
            SocialMediaUrl: socialMedia.SocialMediaUrl,
            UserID:         socialMedia.UserID,
            User: dto.UserMinimalInfoDTO{
                Username: socialMedia.User.Username,
                Email:    socialMedia.User.Email,
            },
            UpdatedAt: socialMedia.UpdatedAt,
            CreatedAt: socialMedia.CreatedAt,
        }
        responseDTOs = append(responseDTOs, responseDTO)
    }

    return responseDTOs, nil
}

func (s *SocialMediaService) EditSocialMedia(userID, socialMediaID uuid.UUID, editDTO dto.EditSocialMediaDTO) (domain.SocialMedia, error) {
	var socialMedia domain.SocialMedia


	tx := s.DB.Begin()


	if err := tx.Preload("User").First(&socialMedia, "id = ? AND user_id = ?", socialMediaID, userID).Error; err != nil {
		tx.Rollback()
		return domain.SocialMedia{}, err
	}


	socialMedia.Name = editDTO.Name
	socialMedia.SocialMediaUrl = editDTO.SocialMediaUrl


	if err := tx.Model(&socialMedia).Select("Name","SocialMediaUrl").Updates(domain.SocialMedia{
		Name: socialMedia.Name,
		SocialMediaUrl: socialMedia.SocialMediaUrl,
	}).Error; err != nil {
		tx.Rollback() 
		return domain.SocialMedia{}, err
	}


	tx.Commit()

	return socialMedia, nil
}

func (s *SocialMediaService) DeleteSocialMedia(userID, socialMediaID uuid.UUID) error {

	tx := s.DB.Begin()

	var socialMedia domain.SocialMedia

	if err := tx.Where("id = ? AND user_id = ?", socialMediaID, userID).First(&socialMedia).Error; err != nil {
		tx.Rollback() 
		return err
	}

	if err := tx.Delete(&socialMedia).Error; err != nil {
		tx.Rollback() 
		return err
	}
	tx.Commit()

	return nil
}



