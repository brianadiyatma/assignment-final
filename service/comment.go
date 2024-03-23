package service

import (
	dto "assignment-final/model/DTO"
	"assignment-final/model/domain"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CommentService struct {
	DB *gorm.DB
}

func NewCommentService(db *gorm.DB) *CommentService {
	return &CommentService{
		DB: db,
	}
}

func (cs *CommentService) CreateComment(userID uuid.UUID, req dto.CreateCommentDTO) (dto.CreateCommentResponseDTO, error) {
    comment := domain.Comment{
        UserID:  userID,
        PhotoID: req.PhotoID,
        Message: req.Message,
    }


    if err := cs.DB.Create(&comment).Error; err != nil {
        return dto.CreateCommentResponseDTO{}, err
    }

    response := dto.CreateCommentResponseDTO{
        ID:        comment.ID,
        Message:   comment.Message,
        PhotoID:   comment.PhotoID,
        UserID:    comment.UserID,
        CreatedAt: comment.CreatedAt,
    }

    return response, nil
}

func (s *CommentService) DeleteComment(commentID uuid.UUID, userID uuid.UUID) error {
    var comment domain.Comment
    if err := s.DB.First(&comment, "id = ?", commentID).Error; err != nil {
        return err
    }

    if comment.UserID != userID {
        return errors.New("Tidak berhak menghapus comment orang lain")
    }

  
    if err := s.DB.Delete(&comment).Error; err != nil {
        return err // Error deleting the comment
    }

    return nil
}