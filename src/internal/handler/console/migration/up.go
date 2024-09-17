package migration

import "github.com/urfave/cli/v2"

func (m *Migration) Up(ctx *cli.Context) error {
	migrator, err := m.createMigrator()
	if err != nil {
		m.Container.Logger.Error(err.Error())
	}

	return nil
}
