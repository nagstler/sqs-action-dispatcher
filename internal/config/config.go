package config

import (
	"errors"
	"os"
)

// Define a Config struct that holds the application's configuration settings
type Config struct {
	SQSQueueURL string
	AWSRegion   string
}

// This function that reads the configuration settings from environment variables or a configuration file, initializes a Config struct, and returns it.
func Load() *Config {
	sqsQueueURL := os.Getenv("SQS_QUEUE_URL")
	if sqsQueueURL == "" {
		panic(errors.New("SQS_QUEUE_URL environment variable is not set"))
	}

	awsRegion := os.Getenv("AWS_REGION")
	if awsRegion == "" {
		panic(errors.New("AWS_REGION environment variable is not set"))
	}

	return &Config{
		SQSQueueURL: sqsQueueURL,
		AWSRegion:   awsRegion,
	}
}

// This function is used by the Load() function to read the value of the SQS_QUEUE_URL environment variable.
func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
