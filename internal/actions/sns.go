package actions

import (
	"encoding/json"
	"fmt"
)

type SNSAction struct{}

type SNSActionData struct {
	TopicArn string `json:"topic_arn"`
	Message  string `json:"message"`
}

func (a *SNSAction) Execute(data json.RawMessage) error {
	// var snsData SNSActionData
	// err := json.Unmarshal(data, &snsData)
	// if err != nil {
	// 	return fmt.Errorf("failed to unmarshal SNS action data: %w", err)
	// }

	// sess := session.Must(session.NewSession(&aws.Config{
	// 	Region: aws.String("your_aws_region"),
	// }))
	// snsClient := sns.New(sess)

	// input := &sns.PublishInput{
	// 	Message:  aws.String(snsData.Message),
	// 	TopicArn: aws.String(snsData.TopicArn),
	// }

	// _, err = snsClient.Publish(input)
	// if err != nil {
	// 	return fmt.Errorf("failed to publish SNS message: %w", err)
	// }

	fmt.Printf("Received message: %s\n", data)

	return nil
}
