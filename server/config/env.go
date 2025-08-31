package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
)

func LoadEnv() error {
	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error getting current working directory: %w", err)
	}

	// Load main .env file
	mainEnvPath := filepath.Join(cwd, ".env")
	if err := godotenv.Load(mainEnvPath); err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}

	mode := os.Getenv("APP_ENV")
	var envFile string

	switch mode {
	case "test":
		envFile = ".env.test"
	case "release":
		envFile = ".env.prod"
	default:
		return nil
	}

	additionalEnvPath := filepath.Join(cwd, envFile)

	// Check if file exists
	if _, err := os.Stat(additionalEnvPath); os.IsNotExist(err) {
		return fmt.Errorf("additional env file does not exist: %s", additionalEnvPath)
	}

	// Overwrites potential duplicate variables
	if err := godotenv.Overload(additionalEnvPath); err != nil {
		return fmt.Errorf("error loading additional env file: %w", err)
	}

	return nil
}
