package pgsql

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresStorage struct {
	db *gorm.DB
}

func NewPostgresStorage(databaseDsn string) (*PostgresStorage, error) {
	db, err := gorm.Open(postgres.Open(databaseDsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &PostgresStorage{db: db}, nil
}
