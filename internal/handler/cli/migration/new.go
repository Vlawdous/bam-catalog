package migration

import (
	"errors"
	"github.com/urfave/cli/v2"
	"os"
	"strconv"
	"time"
)

func (m *Migration) New(c *cli.Context) error {
	name := c.String("name")
	if name == "" {
		return errors.New("name is required")
	}

	newFileName := strconv.FormatInt(time.Now().Unix(), 10) + "_" + name

	var newMigrationPath string = m.Container.Config.AppConfig.MigrationPath + "/"

	newUpMigrationFilePath := newMigrationPath + newFileName + ".up.sql"
	newDownMigrationFilePath := newMigrationPath + newFileName + ".down.sql"

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
