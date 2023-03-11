package app

import (
	"fmt"
	"net/http"

	"github.com/andrsj/feedback-service/internal/delivery/http/handlers"
	"github.com/andrsj/feedback-service/internal/delivery/http/router"
	"github.com/andrsj/feedback-service/internal/delivery/http/server"
	"github.com/andrsj/feedback-service/internal/infrastructure/cache/memory"
	repo "github.com/andrsj/feedback-service/internal/infrastructure/db/memory"
	"github.com/andrsj/feedback-service/internal/services/feedback"
	"github.com/andrsj/feedback-service/pkg/logger"
)

type App struct {
	server *http.Server
	logger logger.Logger
}

func New(logger logger.Logger) (*App, error) {
	logger = logger.Named("app")

	// business logic
	feedbackRepo := repo.New(logger)
	service := feedback.New(feedbackRepo, logger)
	handlers := handlers.New(service, logger)

	// router setting up
	cache := memory.New(logger)
	router := router.New(cache, logger)
	router.Register(handlers)

	server := server.New(router)

	return &App{
		server: server,
		logger: logger,
	}, nil
}

func (a *App) Start() error {
	a.logger.Info("Starting the application", logger.M{"address": a.server.Addr})

	err := a.server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("can't listen and serve: %w", err)
	}

	return nil
}
