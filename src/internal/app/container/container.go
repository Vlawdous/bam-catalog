package container

import (
	appconfig "bam-catalog/internal/app/config"
	"bam-catalog/internal/app/storage/pgsql"
	"log"
)

type Container struct {
	Config          *appconfig.Config
	Logger          *log.Logger
	PostgresStorage *pgsql.PostgresStorage
}

func NewContainer(config *appconfig.Config, logger *log.Logger) (*Container, error) {
	pgsqlStorage, err := pgsql.NewPostgresStorage(config.AppConfig.DatabaseDSN)
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