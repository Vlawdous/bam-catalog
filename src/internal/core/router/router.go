package router

import (
	"bam-catalog/internal/core/container"
	"bam-catalog/internal/handler/http/service"
	"github.com/go-chi/chi/v5"
)

func NewRouter(container *container.Container) (chi.Router, error) {
	router := chi.NewRouter()

	router.Route("/v1", func(r chi.Router) {
		r.Route("/service", func(r chi.Router) {
			serviceHandler := service.NewService(container)

			r.Get("/ping", serviceHandler.Ping)
		})
	})

	return router, nil
}
