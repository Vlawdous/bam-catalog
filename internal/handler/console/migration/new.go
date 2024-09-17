package migration

import (
	"github.com/urfave/cli/v2"
	"os"
	"time"
)

func (m *Migration) New(_ *cli.Context) error {
	newFileName := time.Now().Format("20060102_150405")

	var newMigrationPath string = m.Container.Config.AppConfig.MigrationPath + "/"

	newUpMigrationFilePath := newMigrationPath + newFileName + ".sql.up"
	newDownMigrationFilePath := newMigrationPath + newFileName + ".sql.down"

	newUpMigrationFile, err := os.Create(newUpMigrationFilePath)

	if err != nil {
		return err
	}
	defer newUpMigrationFile.Close()

	newDownMigrationFile, err := os.Create(newDownMigrationFilePath)

	if err != nil {
		return err
	}
	defer newDownMigrationFile.Close()

	return nil
}
