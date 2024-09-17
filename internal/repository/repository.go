package repository

import "bam-catalog/internal/core/storage/gorm"

type Repository struct {
	*gorm.Storage
}

func NewRepository(storage *gorm.Storage) *Repository {
	return &Repository{storage}
}
