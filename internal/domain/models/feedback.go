package models

import "time"

type FeedbackInput struct {
	CustomerName string `json:"customer_name"` //nolint:tagliatelle
	Email        string `json:"email"`
	FeedbackText string `json:"feedback_text"` //nolint:tagliatelle
	Source       string `json:"source"`
}

type Feedback struct {
	ID           string    `json:"id" gorm:"primaryKey"`
	CustomerName string    `json:"customer_name"` //nolint:tagliatelle
	Email        string    `json:"email"`
	FeedbackText string    `json:"feedback_text"` //nolint:tagliatelle
	Source       string    `json:"source"`
	CreatedAt    time.Time `json:"-" gorm:"created_at"`
	UpdatedAt    time.Time `json:"-" gorm:"updated_at"`
}
