package repositories

import "github.com/andrsj/feedback-service/internal/domain/models"

type FeedbackRepository interface {
	Create(feedback *models.FeedbackInput) (feedbackID string, err error)
	GetByID(feedbackID string) (feedback *models.Feedback, err error)
	GetAll() (feedbacks []*models.Feedback, err error)
}