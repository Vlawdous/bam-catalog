package container

import (
	appconfig "bam-catalog/internal/core/config"
	"bam-catalog/internal/core/storage/gorm"
	"log"
)

type Container struct {
	Config          *appconfig.Config
	Logger          *log.Logger
	PostgresStorage *gorm.Storage
}

func NewContainer(config *appconfig.Config, logger *log.Logger) (*Container, error) {
	pgsqlStorage, err := gorm.NewPostgresStorage(config.AppConfig.DatabaseDSN)
	if err != nil {
		return nil, err
	}

	container := Container{
		config,
		logger,
		pgsqlStorage,
	}

	return &container, nil
}
