package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
)

var projectPath string

func init() {
	execPath, err := filepath.Abs(filepath.Dir(os.Args[0]))

	if err != nil {
		log.Fatal("Не удалось найти директорию проекта: " + err.Error())
	}

	projectPath = execPath

	localEnvErr := godotenv.Load(filepath.Join(projectPath, ".env.local"))
	if localEnvErr != nil {
		log.Print("Не удалось прочитать переменные окружения в файле .env.local: " + localEnvErr.Error())
	}

	envErr := godotenv.Load(filepath.Join(projectPath, ".env"))
	if envErr != nil {
		log.Print("Не удалось прочитать переменные окружения в файле .env: " + envErr.Error())
	}

	if localEnvErr != nil && envErr != nil {
		log.Fatal("Не удалось прочитать ни одной переменной окружения")
	}
}

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
	ProjectPath   string
	LogFilePath   string `env:"LOG_FILE_PATH"`
	MigrationPath string `env:"MIGRATION_PATH"`
	DatabaseDSN   string `env:"DATABASE_DSN,required"`
}

func GetConfig() (*Config, error) {
	config := &Config{}
	err := env.Parse(config)

	if err != nil {
		return nil, err
	}

	config.AppConfig.ProjectPath = projectPath
	loadEmptyOptionalProperties(config)

	return config, nil
}

func loadEmptyOptionalProperties(config *Config) {
	if "" == config.AppConfig.LogFilePath {
		config.AppConfig.LogFilePath = config.AppConfig.ProjectPath + "/logs"
	}

	if "" == config.AppConfig.MigrationPath {
		config.AppConfig.MigrationPath = config.AppConfig.ProjectPath + "/migrations"
	}
}
