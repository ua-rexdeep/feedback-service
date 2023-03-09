package handlers

import (
	"fmt"
	"net/http"

	"github.com/andrsj/feedback-service/pkg/logger"
)

type Handlers interface {
	Status(w http.ResponseWriter, r *http.Request)
}

func New(logger logger.Logger) *handlers {
	return &handlers{
		logger: logger.Named("handlers"),
	}
}

type handlers struct {
	logger logger.Logger
}

func (h *handlers) Status(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Status gotted", nil)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Ok")
}
