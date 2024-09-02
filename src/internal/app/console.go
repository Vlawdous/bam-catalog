package app

import (
	appconfig "bam-catalog/internal/app/config"
	appcontainer "bam-catalog/internal/app/container"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

type ConsoleHandlerFunc func(io io.ReadWriter) error

func StartCommand() error {
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

	//TODO: доделать
	fmt.Print(container)

	return nil
}
