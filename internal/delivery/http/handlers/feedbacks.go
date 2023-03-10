package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/andrsj/feedback-service/pkg/logger"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type Feedback struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}

type Feedbacks []*Feedback

var (
	errJSONcoding = errors.New("json en/de-coding error")
	errIDParamIsMissing  = errors.New("id parameter is missing")
	errIDIsMissing  = errors.New("id is missing")
)

// TODO: DELETE THIS SHIT!
//nolint 
var feedbacks = Feedbacks{
	{
		ID:      generateID(),
		Message: "Example feedback 1",
	},
	{
		ID:      generateID(),
		Message: "Example feedback 2",
	},
}

// GetFeedback GET /feedback/{id}.
func (h *handlers) GetFeedback(w http.ResponseWriter, r *http.Request) {
	feedbackID := chi.URLParam(r, "id")
	if feedbackID == "" {
		err := errIDParamIsMissing
		h.logger.Error(err.Error(), logger.M{"err": err})
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	var feedback *Feedback

	for k, v := range feedbacks {
		if v.ID == feedbackID {
			feedback = feedbacks[k]
		}
	}

	if feedback == nil {
		err := errIDIsMissing
		h.logger.Error(err.Error(), logger.M{"err": err})
		http.Error(w, err.Error(), http.StatusNotFound)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	err := json.NewEncoder(w).Encode(feedback)
	if err != nil {
		err = errJSONcoding
		h.logger.Error(err.Error(), logger.M{"err": err})
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}

// GetAllFeedback GET /feedbacks.
func (h *handlers) GetAllFeedback(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(feedbacks)
	if err != nil {
		err = errJSONcoding
		h.logger.Error(err.Error(), logger.M{"err": err})
		http.Error(w, errJSONcoding.Error(), http.StatusInternalServerError)

		return
	}
}

// CreateFeedback POST /feedback.
func (h *handlers) CreateFeedback(w http.ResponseWriter, r *http.Request) {
	var feedback Feedback

	err := json.NewDecoder(r.Body).Decode(&feedback)
	if err != nil {
		err = errJSONcoding
		h.logger.Error(err.Error(), logger.M{"err": err})
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	feedback.ID = generateID()

	feedbacks = append(feedbacks, &feedback)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(map[string]string{"id": feedback.ID})
	if err != nil {
		err = errJSONcoding
		h.logger.Error(err.Error(), logger.M{"err": err})
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}

func generateID() string {
	return uuid.NewString()
}