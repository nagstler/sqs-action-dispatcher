package utils

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/pkg/errors"
)

// ParseJSON unmarshals a JSON string into the given interface.
func ParseJSON(data string, v interface{}) error {
	err := json.Unmarshal([]byte(data), v)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal JSON")
	}
	return nil
}

// LogError logs the given error if it is not nil.
func LogError(err error) {
	if err != nil {
		log.Printf("Error: %v", err)
	}
}

// NewAWSSession creates a new AWS session with the given region.
func NewAWSSession(region string) (*session.Session, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create AWS session")
	}
	return sess, nil
}

// IsAWSError checks if the given error is of the specified AWS error code.
func IsAWSError(err error, code string) bool {
	if awsErr, ok := err.(awserr.Error); ok {
		return strings.Contains(awsErr.Code(), code)
	}
	return false
}
