package poller

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
)

// Message represents an SQS message with Body and ReceiptHandle.
type Message struct {
	Body          string
	ReceiptHandle *string
}

// Poller is responsible for polling messages from the SQS queue.
type Poller struct {
	sqsClient sqsiface.SQSAPI
	queueURL  string
}

// NewPoller creates a new Poller instance with the given SQS client and queue URL.
func NewPoller(sqsClient sqsiface.SQSAPI, queueURL string) *Poller {
	return &Poller{
		sqsClient: sqsClient,
		queueURL:  queueURL,
	}
}

// Poll fetches messages from the SQS queue and returns them as a slice of Message structs.
func (p *Poller) Poll() ([]Message, error) {
	input := &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(p.queueURL),
		MaxNumberOfMessages: aws.Int64(10),
		WaitTimeSeconds:     aws.Int64(20),
	}

	output, err := p.sqsClient.ReceiveMessage(input)
	if err != nil {
		return nil, err
	}

	messages := make([]Message, len(output.Messages))
	for i, msg := range output.Messages {
		messages[i] = Message{
			Body:          aws.StringValue(msg.Body),
			ReceiptHandle: msg.ReceiptHandle,
		}
	}

	return messages, nil
}

// Send messages to SQS - Dead Letter Queue
func (p *Poller) SendMessageToDLQ(messageBody string, dlqURL string) error {
	_, err := p.sqsClient.SendMessage(&sqs.SendMessageInput{
		MessageBody: aws.String(messageBody),
		QueueUrl:    aws.String(dlqURL),
	})
	return err
}

// DeleteMessage deletes a message from the SQS queue using the message's ReceiptHandle.
func (p *Poller) DeleteMessage(receiptHandle *string) error {
	_, err := p.sqsClient.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      aws.String(p.queueURL),
		ReceiptHandle: receiptHandle,
	})
	return err
}
