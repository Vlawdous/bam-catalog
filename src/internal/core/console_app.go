package core

import (
	appconfig "bam-catalog/internal/config"
	config2 "bam-catalog/internal/core/config"
	appcontainer "bam-catalog/internal/core/container"
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
	config := config2.GetConfig()

	log.Print("Initialization container")
	container, err := appcontainer.NewContainer(&config, &log.Logger{})
	if err != nil {
		return err
	}

	//TODO: доделать
	fmt.Print(container)

	return nil
}
