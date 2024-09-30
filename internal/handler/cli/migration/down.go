package migration

import (
	"errors"
	"github.com/urfave/cli/v2"
)

func (m *Migration) Down(c *cli.Context) error {
	migrator, err := m.createMigrator()
	if err != nil {
		return err
	}

	if migrator == nil {
		return errors.New("empty migrator")
	}

	steps := c.Int("steps")

	if steps != 0 {
		err = migrator.Steps(steps)
	} else {
		err = migrator.Up()
	}

	if err != nil {
		return err
	}

	return nil
}
