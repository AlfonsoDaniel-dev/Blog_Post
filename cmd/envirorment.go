package main

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
	"strings"
)

var envVars = []string{"DB_HOST", "DB_PORT", "DB_USER", "DB_PASSWORD", "DB_NAME", "APP_HOST", "APP_PORT", "ALLOWED_ORIGINS"}

func loadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	return nil
}

func ValidateEnvVars() error {
	for _, envVar := range envVars {
		if strings.TrimSpace(os.Getenv(envVar)) == "" {
			return errors.New(envVar + " is required")
		}
	}
	return nil
}

/*
func ValidateEnvVar() error {
	if strings.TrimSpace(os.Getenv("DB_HOST")) == "" {
		return errors.New("DB_HOST environment variable not set")
	}

	if strings.TrimSpace(os.Getenv("DB_HOST")) == "" {
		return errors.New("DB_HOST environment variable not set")
	}

} */
