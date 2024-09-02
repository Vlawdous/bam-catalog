package config

import (
	"fmt"
	"github.com/caarlos0/env/v11"
	"log"
	"os"
	"path/filepath"
	"sync"
)

type Config struct {
	Env          string       `env:"ENV,required"`
	ListenConfig ListenConfig `envPrefix:"LISTEN_"`
	AppConfig    AppConfig    `envPrefix:"APP_"`
}

type ListenConfig struct {
	Address             string `env:"ADDRESS" envDefault:"0.0.0.0"`
	Port                string `env:"PORT" envDefault:"8080"`
	ReadTimeoutSeconds  int    `env:"READ_TIMEOUT" envDefault:"10"`
	WriteTimeoutSeconds int    `env:"WRITE_TIMEOUT" envDefault:"10"`
	IdleTimeoutSeconds  int    `env:"IDLE_TIMEOUT" envDefault:"10"`
}

type AppConfig struct {
	LogFilePath string `env:"LOG_FILE_PATH"`
	DatabaseDSN string `env:"DATABASE_DSN,required"`
}

var (
	config = Config{}
	once   sync.Once
)

func GetConfig() Config {
	once.Do(func() {
		err := env.Parse(&config)

		if err != nil {
			log.Fatal(err)
		}

		loadEmptyOptionalProperties()
	})

	return config
}

func loadEmptyOptionalProperties() {
	if "" == config.AppConfig.LogFilePath {
		config.AppConfig.LogFilePath = fmt.Sprintf("%s/logs", filepath.Dir(os.Args[0]))
	}
}
