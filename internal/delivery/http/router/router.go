package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/andrsj/feedback-service/internal/delivery/http/handlers"
	"github.com/andrsj/feedback-service/internal/delivery/http/middlewares"
	"github.com/andrsj/feedback-service/internal/infrastructure/cache"
	"github.com/andrsj/feedback-service/pkg/logger"
)

type Router struct {
	router *chi.Mux
	logger logger.Logger

	cacheMiddleware func(next http.Handler) http.Handler
}

func New(cache cache.Cache, logger logger.Logger) *Router {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	cacheMiddleware := middlewares.CacheMiddleware(cache)

	return &Router{
		router: router,
		logger: logger.Named("router"),
		cacheMiddleware: cacheMiddleware,
	}
}

func (r *Router) Register(handler handlers.Handlers) {
	r.logger.Info("Registring handlers", nil)
	r.router.Get("/", handler.Status)

	r.router.Group(
		func(router chi.Router) {
			router.Use(r.cacheMiddleware)
			router.Get("/feedbacks", handler.GetAllFeedback)
			router.Get("/feedback/{id}", handler.GetFeedback)
		},
	)
	r.router.Post("/feedback", handler.CreateFeedback)

	err := chi.Walk(
		r.router, 
		func(method string, route string, _ http.Handler, _ ...func(http.Handler) http.Handler) error {
        	r.logger.Info(fmt.Sprintf("%-5s -> %s", method, route), nil)

        	return nil
    	},
	)
	if err != nil {
		panic(err)
	}
}

func (r *Router) GetChiMux() *chi.Mux {
	return r.router
}
