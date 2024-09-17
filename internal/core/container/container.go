package container

import (
	appConfig "bam-catalog/internal/core/config"
	"bam-catalog/internal/core/logger"
	"bam-catalog/internal/core/storage/gorm"
	"bam-catalog/internal/repository"
)

type Container struct {
	Config       *appConfig.Config
	Logger       *logger.Logger
	Repositories *Repositories
}

type Repositories struct {
	BookRepository *repository.BookRepository
}

func NewContainer(config *appConfig.Config, logger *logger.Logger) (*Container, error) {
	pgsqlStorage, err := gorm.NewPostgresStorage(config.AppConfig.DatabaseDSN)
	if err != nil {
		return nil, err
	}

	baseRepository := &repository.Repository{Storage: pgsqlStorage}

	repositories := &Repositories{
		BookRepository: &repository.BookRepository{Repository: baseRepository},
	}

	container := Container{
		config,
		logger,
		repositories,
	}

	return &container, nil
}
