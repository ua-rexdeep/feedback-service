package gorm

import (
	"fmt"
	"time"

	"gorm.io/gorm"
	"github.com/google/uuid"

	"github.com/andrsj/feedback-service/internal/domain/models"
	"github.com/andrsj/feedback-service/internal/domain/repositories"
	"github.com/andrsj/feedback-service/pkg/logger"

)

type FeedbackRepository struct {
	db     *gorm.DB
	logger logger.Logger
}

var _ repositories.FeedbackRepository = (*FeedbackRepository)(nil)

func NewFeedbackRepository(db *gorm.DB, logger logger.Logger) *FeedbackRepository {
	return &FeedbackRepository{
		db:     db,
		logger: logger.Named("gormORM"),
	}
}

func (r *FeedbackRepository) Create(*models.FeedbackInput) (string, error) {
	r.logger.Info("Creating 'Feedback'", nil)

	var (
		feedback   models.Feedback
		feedbackID = uuid.NewString()
	)

	feedback = models.Feedback{
		ID:           feedbackID,
		CustomerName: feedback.CustomerName,
		Email:        feedback.Email,
		FeedbackText: feedback.FeedbackText,
		Source:       feedback.Source,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	err := r.db.Create(&feedback).Error
	if err != nil {
		r.logger.Error("Failed to create feedback into DB", logger.M{"err": err})

		return "", fmt.Errorf("failed to create feedback into DB: %w", err)
	}

	r.logger.Info("Feedback created successfully", logger.M{"id": feedbackID})

	return feedbackID, nil
}

func (r *FeedbackRepository) GetByID(feedbackID string) (*models.Feedback, error) {
	var feedback models.Feedback

	r.logger.Info("Getting 'Feedback' by ID", logger.M{
		"feedbackID": feedbackID,
	})

	err := r.db.First(&feedback, feedbackID).Error
	if err != nil {
		r.logger.Error("Failed to get feedback from DB", logger.M{
			"feedbackID": feedbackID,
			"error":      err.Error(),
		})

		return nil, fmt.Errorf("failed to get feedback from DB: %w", err)
	}

	r.logger.Info("Got 'Feedback' by ID", logger.M{"feedbackID": feedbackID})

	return &feedback, nil
}

func (r *FeedbackRepository) GetAll() ([]*models.Feedback, error) {
	var feedbacks []*models.Feedback

	r.logger.Info("Get all 'Feedback's", nil)

	err := r.db.Order("created_at").Find(&feedbacks).Error
	if err != nil {
		r.logger.Error("Failed to get feedback from DB", logger.M{"error": err.Error()})

		return nil, fmt.Errorf("failed to get feedbacks from DB: %w", err)
	}

	r.logger.Info("Got all 'Feedback's", logger.M{"count": len(feedbacks)})

	return feedbacks, nil
}
