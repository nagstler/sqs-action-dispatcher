package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/nagstler/sqs-action-dispatcher/internal/config"
	"github.com/nagstler/sqs-action-dispatcher/internal/dispatcher"
	"github.com/nagstler/sqs-action-dispatcher/internal/poller"
)

const numberOfWorkers = 5

// worker function processes messages in a concurrent fashion using goroutines.
// It polls messages from the SQS queue, dispatches actions, and deletes messages after successful processing.
func worker(id int, wg *sync.WaitGroup, sqsPoller *poller.Poller, actionDispatcher *dispatcher.Dispatcher, dlqURL string) {
	defer wg.Done()

	for {
		messages, err := sqsPoller.Poll()
		if err != nil {
			log.Printf("Error polling messages by worker %d: %v", id, err)
			continue
		}

		for _, msg := range messages {
			fmt.Printf("Worker %d received message: %s\n", id, msg.Body)

			err := actionDispatcher.Dispatch(msg.Body)
			if err != nil {
				log.Printf("Error dispatching message by worker %d: %v", id, err)

				// Send the message to DLQ and delete the message from the source queue
				if err := sqsPoller.SendMessageToDLQ(msg.Body, dlqURL); err != nil {
					log.Printf("Error sending message to DLQ by worker %d: %v", id, err)
				}
			}

			// Delete the message from the source queue after processing (success or failure)
			if err := sqsPoller.DeleteMessage(msg.ReceiptHandle); err != nil {
				log.Printf("Error deleting message by worker %d: %v", id, err)
			}
		}
	}
}

func main() {
	// Load the configuration
	cfg := config.Load()

	// Create a new AWS session
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(cfg.AWSRegion),
	}))

	// Create a new SQS client and poller instance
	sqsClient := sqs.New(sess)
	sqsPoller := poller.NewPoller(sqsClient, cfg.SQSQueueURL)

	// Initialize the action dispatcher
	actionDispatcher := dispatcher.NewDispatcher()

	// Create a wait group to synchronize the worker goroutines
	var wg sync.WaitGroup

	dlqURL := cfg.SQSDLQURL
	// Start the worker goroutines
	for i := 1; i <= numberOfWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg, sqsPoller, actionDispatcher, dlqURL)
	}

	// Wait for all worker goroutines to complete
	wg.Wait()
}
