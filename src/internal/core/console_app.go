package core

import (
	appConfig "bam-catalog/internal/core/config"
	appContainer "bam-catalog/internal/core/container"
	appLogger "bam-catalog/internal/core/logger"
	"bam-catalog/internal/handler"
	"bam-catalog/internal/handler/console/migration"
	"github.com/urfave/cli/v2"
	"log"
)

func RunCli() error {
	log.Print("Инициализация и получение конфига")
	config, err := appConfig.GetConfig()
	if err != nil {
		return err
	}

	log.Print("Инициализация логгера")
	logger := appLogger.NewLogger(config.AppConfig.LogFilePath, config.Env)

	log.Print("Инициализация контейнера")
	container, err := appContainer.NewContainer(config, logger)
	if err != nil {
		return err
	}

	consoleHandler := &handler.Handler{Container: container}

	migrationHandler := &migration.Migration{Handler: consoleHandler}

	cliApp := cli.App{}
}

func getCommands()
