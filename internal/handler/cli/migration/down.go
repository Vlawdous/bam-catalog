package migration

import (
	"errors"
	"github.com/urfave/cli/v2"
)

func (m *Migration) Down(_ *cli.Context) error {
	migrator, err := m.createMigrator()
	if err != nil {
		return err
	}

	if migrator == nil {
		return errors.New("empty migrator")
	}

	err = migrator.Down()
	if err != nil {
		return err
	}

	return nil
}
