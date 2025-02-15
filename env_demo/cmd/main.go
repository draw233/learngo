package main

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func init() {
	err := godotenv.Load("default.env", ".env")
	if err != nil {
		slog.Info("Error loading .env file, using default values")
	}
}

func main() {
	shell := getEnv("SHELL", "")
	xx := getEnv("ENV_FILE", "default")
	slog.Info(shell)
	slog.Info(xx)
}
