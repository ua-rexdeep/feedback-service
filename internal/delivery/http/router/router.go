package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/andrsj/feedback-service/internal/delivery/http/handlers"
	"github.com/andrsj/feedback-service/pkg/logger"

)

type Router struct {
	router *chi.Mux
	logger logger.Logger
}

func New(logger logger.Logger) *Router {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	return &Router{
		router: router,
		logger: logger.Named("router"),
	}
}

func (r *Router) Register(handler handlers.Handlers) {
	r.logger.Info("Registring handlers", nil)
	r.router.Get("/", handler.Status)
}

func (r *Router) GetChiMux() *chi.Mux {
	return r.router
}
