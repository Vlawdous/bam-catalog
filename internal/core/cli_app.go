package core

import (
	appConfig "bam-catalog/internal/core/config"
	appContainer "bam-catalog/internal/core/container"
	appLogger "bam-catalog/internal/core/logger"
	"bam-catalog/internal/handler"
	"bam-catalog/internal/handler/cli/migration"
	"github.com/urfave/cli/v2"
	"os"
)

func RunCli() error {
	config, err := appConfig.GetConfig()
	if err != nil {
		return err
	}

	logger := appLogger.NewLogger(config.AppConfig.LogFilePath, config.Env)

	container, err := appContainer.NewContainer(config, logger)
	if err != nil {
		return err
	}

	cliApp := cli.App{Commands: getCommands(container)}

	err = cliApp.Run(os.Args)
	if err != nil {
		return err
	}

	return nil
}

func getCommands(container *appContainer.Container) []*cli.Command {
	consoleHandler := &handler.Handler{Container: container}

	migrationHandler := &migration.Migration{Handler: consoleHandler}

	return []*cli.Command{
		{
			Name:    "migrate",
			Aliases: []string{"m"},
			Usage:   "Общая команда миграций",
			Subcommands: []*cli.Command{
				{
					Name:   "up",
					Usage:  "Поднять версионность БД",
					Action: migrationHandler.Up,
					Flags: []cli.Flag{
						&cli.IntFlag{
							Name:    "steps",
							Aliases: []string{"s"},
						},
					},
				},
				{
					Name:   "down",
					Usage:  "Опустить версионность БД",
					Action: migrationHandler.Down,
					Flags: []cli.Flag{
						&cli.IntFlag{
							Name:    "steps",
							Aliases: []string{"s"},
						},
					},
				},
				{
					Name:   "new",
					Usage:  "Создать новую миграцию",
					Action: migrationHandler.New,
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:     "name",
							Aliases:  []string{"n"},
							Required: true,
						},
					},
				},
			},
		},
	}
}
