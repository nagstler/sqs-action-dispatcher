package main

import (
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/nagstler/sqs-action-dispatcher/internal/config"
	"github.com/nagstler/sqs-action-dispatcher/internal/dispatcher"
	"github.com/nagstler/sqs-action-dispatcher/internal/poller"
)

func main() {
	cfg := config.Load()

	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(cfg.AWSRegion),
	}))

	sqsClient := sqs.New(sess)
	sqsPoller := poller.NewPoller(sqsClient, cfg.SQSQueueURL)

	actionDispatcher := dispatcher.NewDispatcher()

	for {
		messages, err := sqsPoller.Poll()
		if err != nil {
			log.Printf("Error polling messages: %v", err)
		}

		for _, msg := range messages {
			fmt.Printf("Received message: %s\n", msg.Body)

			if err := actionDispatcher.Dispatch(msg.Body); err != nil {
				log.Printf("Error dispatching message: %v", err)
			}
		}

		time.Sleep(5 * time.Second)
	}
}
