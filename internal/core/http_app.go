package core

import (
	appConfig "bam-catalog/internal/core/config"
	appContainer "bam-catalog/internal/core/container"
	appLogger "bam-catalog/internal/core/logger"
	appRouter "bam-catalog/internal/core/router"
	"fmt"
	"log"
	"net/http"
	"time"
)

func RunHttp() error {
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

	log.Print("Инициализация HTTP-сервера")
	server, err := createHttpServer(container)
	if err != nil {
		return err
	}

	log.Print("Запуск HTTP-сервера")
	err = server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

func createHttpServer(container *appContainer.Container) (*http.Server, error) {
	router, err := appRouter.NewRouter(container)
	if err != nil {
		return nil, err
	}

	addr := fmt.Sprintf("%s:%s", container.Config.ListenConfig.Address, container.Config.ListenConfig.Port)
	server := &http.Server{
		Addr:         addr,
		Handler:      router,
		ReadTimeout:  time.Duration(container.Config.ListenConfig.ReadTimeoutSeconds) * time.Second,
		WriteTimeout: time.Duration(container.Config.ListenConfig.WriteTimeoutSeconds) * time.Second,
		IdleTimeout:  time.Duration(container.Config.ListenConfig.IdleTimeoutSeconds) * time.Second,
	}

	return server, nil
}
