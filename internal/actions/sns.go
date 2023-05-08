package actions

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

type SNSAction struct{}

type SNSActionData struct {
	TopicArn string `json:"topic_arn"`
	Message  string `json:"message"`
}

func (a *SNSAction) Execute(data json.RawMessage) error {
	var snsData SNSActionData
	err := json.Unmarshal(data, &snsData)
	if err != nil {
		return fmt.Errorf("failed to unmarshal SNS action data: %w", err)
	}
	sess := session.Must(session.NewSession(&aws.Config{}))
	snsClient := sns.New(sess)

	input := &sns.PublishInput{
		Message:  aws.String(snsData.Message),
		TopicArn: aws.String(snsData.TopicArn),
	}

	_, err = snsClient.Publish(input)
	if err != nil {
		return fmt.Errorf("failed to publish SNS message: %w", err)
	}

	return nil
}
