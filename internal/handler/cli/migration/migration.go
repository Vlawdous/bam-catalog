package migration

import (
	"bam-catalog/internal/handler"
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Migration struct {
	*handler.Handler
}

func (m *Migration) createMigrator() (*migrate.Migrate, error) {
	appConfig := &m.Container.Config.AppConfig

	db, err := sql.Open("postgres", appConfig.DatabaseDSN)
	if err != nil {
		return nil, err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	return migrate.NewWithDatabaseInstance("file://"+appConfig.MigrationPath, "postgres", driver)
}
