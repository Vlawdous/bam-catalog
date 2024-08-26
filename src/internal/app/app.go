package app

import (
	appconfig "bam-catalog/internal/app/config"
	appcontainer "bam-catalog/internal/app/container"
	approuter "bam-catalog/internal/app/router"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func Run() error {
	log.Print("Collect project path")
	projectPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return err
	}

	log.Print("Initialization dotEnv files")
	if err := appconfig.LoadDotEnv(projectPath); err != nil {
		return err
	}

	log.Print("Collect config properties")
	config := appconfig.GetConfig()

	log.Print("Initialization container")
	container, err := appcontainer.NewContainer(&config, &log.Logger{})
	if err != nil {
		return err
	}

	// TODO: придумать, как его запихать в обработчик запроса
	fmt.Println(container)

	log.Print("Initialization router")
	router, err := approuter.NewRouter()
	if err != nil {
		return err
	}

	log.Print("Initialization HTTP server")
	addr := fmt.Sprintf("%s:%s", config.ListenConfig.Address, config.ListenConfig.Port)
	server := &http.Server{
		Addr:         addr,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Print("Start HTTP server")
	err = server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
