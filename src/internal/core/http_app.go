package core

import (
	appConfig "bam-catalog/internal/core/config"
	appContainer "bam-catalog/internal/core/container"
	appRouter "bam-catalog/internal/core/router"
	"fmt"
	"log"
	"net/http"
	"time"
)

func Run() error {
	log.Print("Инициализация и получение конфига")
	config, err := appConfig.GetConfig()

	log.Print("Инициализация контейнера")
	container, err := appContainer.NewContainer(config, &log.Logger{})
	if err != nil {
		return err
	}

	log.Print("Инициализация роутера")
	router, err := appRouter.NewRouter(container)
	if err != nil {
		return err
	}

	log.Print("Инициализация HTTP-сервера")
	addr := fmt.Sprintf("%s:%s", config.ListenConfig.Address, config.ListenConfig.Port)
	server := &http.Server{
		Addr:         addr,
		Handler:      router,
		ReadTimeout:  time.Duration(config.ListenConfig.ReadTimeoutSeconds) * time.Second,
		WriteTimeout: time.Duration(config.ListenConfig.WriteTimeoutSeconds) * time.Second,
		IdleTimeout:  time.Duration(config.ListenConfig.IdleTimeoutSeconds) * time.Second,
	}

	log.Print("Запуск HTTP-сервера")
	err = server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
