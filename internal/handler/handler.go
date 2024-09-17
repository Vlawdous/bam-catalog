package handler

import (
	"bam-catalog/internal/core/container"
)

type Handler struct {
	Container *container.Container
}

func NewHandler(container *container.Container) *Handler {
	return &Handler{
		container,
	}
}
