package core

import (
	appconfig "bam-catalog/internal/core/config"
	appcontainer "bam-catalog/internal/core/container"
	"fmt"
	"io"
	"log"
)

type ConsoleHandlerFunc func(io io.ReadWriter) error

func StartCommand() error {
	log.Print("Collect config properties")
	config, err := appconfig.GetConfig()

	log.Print("Initialization container")
	container, err := appcontainer.NewContainer(config, &log.Logger{})
	if err != nil {
		return err
	}

	//TODO: доделать
	fmt.Print(container)

	return nil
}
