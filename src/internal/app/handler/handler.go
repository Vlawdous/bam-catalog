package handler

import "bam-catalog/internal/app/container"

type Handler struct {
	Container *container.Container
}

func NewHandler(container *container.Container) *Handler {
	return &Handler{
		container,
	}
}
