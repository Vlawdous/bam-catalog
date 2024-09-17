package service

import (
	"bam-catalog/internal/core/container"
	"bam-catalog/internal/handler"
)

type Service struct {
	*handler.Handler
}

func NewService(container *container.Container) *Service {
	return &Service{
		handler.NewHandler(container),
	}
}
