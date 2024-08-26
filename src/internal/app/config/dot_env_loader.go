package config

import (
	"github.com/joho/godotenv"
	"path/filepath"
)

func LoadDotEnv(dotEnvDirectory string) error {
	if err := godotenv.Load(filepath.Join(dotEnvDirectory, ".env.local")); err != nil {
		return err
	}

	if err := godotenv.Load(filepath.Join(dotEnvDirectory, ".env")); err != nil {
		return err
	}

	return nil
}
