package service

import (
	"bam-catalog/internal/app/container"
	"bam-catalog/internal/app/handler"
)

type Service struct {
	*handler.Handler
}

func NewService(container *container.Container) *Service {
	return &Service{
		handler.NewHandler(container),
	}
}
