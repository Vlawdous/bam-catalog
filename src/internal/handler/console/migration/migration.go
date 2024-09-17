package migration

import (
	"bam-catalog/internal/handler"
	"github.com/golang-migrate/migrate/v4"
)

type Migration struct {
	*handler.Handler
}

func (m *Migration) createMigrator() (*migrate.Migrate, error) {
	appConfig := &m.Container.Config.AppConfig
	return migrate.New(appConfig.MigrationPath, appConfig.DatabaseDSN)
}
