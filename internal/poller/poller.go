package poller

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
)

type Message struct {
	Body          string
	ReceiptHandle *string
}

type Poller struct {
	sqsClient sqsiface.SQSAPI
	queueURL  string
}

func NewPoller(sqsClient sqsiface.SQSAPI, queueURL string) *Poller {
	return &Poller{
		sqsClient: sqsClient,
		queueURL:  queueURL,
	}
}

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
