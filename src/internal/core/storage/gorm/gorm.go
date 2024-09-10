package gorm

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	Db *gorm.DB
}

func NewPostgresStorage(databaseDsn string) (*Storage, error) {
	db, err := gorm.Open(postgres.Open(databaseDsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Storage{Db: db}, nil
}
