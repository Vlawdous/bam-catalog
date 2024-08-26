package router

import (
	"bam-catalog/internal/app/handler/service"
	"github.com/go-chi/chi/v5"
)

func NewRouter() (chi.Router, error) {
	router := chi.NewRouter()

	router.Route("/v1", func(r chi.Router) {
		r.Route("/service", func(r chi.Router) {
			r.Get("/ping", service.Ping)
		})
	})

	return router, nil
}
