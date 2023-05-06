package config

import (
	"log"
	"os"
)

// Define a Config struct that holds the application's configuration settings
type Config struct {
	SQSQueueURL string
}

// This function that reads the configuration settings from environment variables or a configuration file, initializes a Config struct, and returns it.
func Load() Config {
	cfg := Config{
		SQSQueueURL: getEnv("SQS_QUEUE_URL", ""),
	}

	if cfg.SQSQueueURL == "" {
		log.Fatal("SQS_QUEUE_URL is required")
	}

	return cfg
}

// This function is used by the Load() function to read the value of the SQS_QUEUE_URL environment variable.
func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
