package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/andrsj/feedback-service/internal/delivery/http/handlers"
	"github.com/andrsj/feedback-service/internal/delivery/http/middlewares"
	"github.com/andrsj/feedback-service/internal/infrastructure/cache/memory"
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

	cacheMiddleware := middlewares.CacheMiddleware(memory.New(logger))
	router.Use(cacheMiddleware)
	
	customMiddleware := middlewares.Mid(10)
	router.Use(customMiddleware)

	return &Router{
		router: router,
		logger: logger.Named("router"),
	}
}

func (r *Router) Register(handler handlers.Handlers) {
	r.logger.Info("Registring handlers", nil)
	r.router.Get("/", handler.Status)
	r.router.Get("/feedback", handler.GetAllFeedback)
	r.router.Get("/feedback/{id}", handler.GetFeedback)
	r.router.Post("/feedback", handler.CreateFeedback)
}

func (r *Router) GetChiMux() *chi.Mux {
	return r.router
}
