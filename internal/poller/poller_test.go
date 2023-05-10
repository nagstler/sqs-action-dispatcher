package poller_test

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
	"github.com/nagstler/sqs-action-dispatcher/internal/poller"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// mockSQSClient is a mock implementation of sqsiface.SQSAPI for testing.
type mockSQSClient struct {
	mock.Mock
	sqsiface.SQSAPI
}

// SendMessage is a mock implementation of sqsiface.SQSAPI.SendMessage.
func (m *mockSQSClient) SendMessage(input *sqs.SendMessageInput) (*sqs.SendMessageOutput, error) {
	args := m.Called(input)
	return args.Get(0).(*sqs.SendMessageOutput), args.Error(1)
}

// DeleteMessage is a mock implementation of sqsiface.SQSAPI.DeleteMessage.
func (m *mockSQSClient) DeleteMessage(input *sqs.DeleteMessageInput) (*sqs.DeleteMessageOutput, error) {
	args := m.Called(input)
	return args.Get(0).(*sqs.DeleteMessageOutput), args.Error(1)
}

// ReceiveMessage is a mock implementation of sqsiface.SQSAPI.ReceiveMessage.
// It returns two dummy messages for testing.
func (m *mockSQSClient) ReceiveMessage(input *sqs.ReceiveMessageInput) (*sqs.ReceiveMessageOutput, error) {
	return &sqs.ReceiveMessageOutput{
		Messages: []*sqs.Message{
			{
				Body:          aws.String("message 1"),
				ReceiptHandle: aws.String("receipt 1"),
			},
			{
				Body:          aws.String("message 2"),
				ReceiptHandle: aws.String("receipt 2"),
			},
		},
	}, nil
}

// TestPoll tests the Poll method of the Poller.
func TestPoll(t *testing.T) {
	// Create a new Poller with the mock SQS client
	p := poller.NewPoller(&mockSQSClient{}, "testQueueURL")

	// Call Poll and verify it returns the correct messages
	messages, err := p.Poll()

	// Assert there's no error and the returned messages are correct
	assert.Nil(t, err)
	assert.Equal(t, 2, len(messages))
	assert.Equal(t, "message 1", messages[0].Body)
	assert.Equal(t, "receipt 1", *messages[0].ReceiptHandle)
	assert.Equal(t, "message 2", messages[1].Body)
	assert.Equal(t, "receipt 2", *messages[1].ReceiptHandle)
}

// TestSendMessageToDLQ tests the SendMessageToDLQ method of the Poller.
func TestSendMessageToDLQ(t *testing.T) {
	// Create a new mock SQS client and a Poller with this client
	mockClient := new(mockSQSClient)
	p := poller.NewPoller(mockClient, "testQueueURL")

	// Set up the expected input and output for the SendMessage call
	messageBody := "Test Message"
	dlqURL := "testDLQURL"

	// Set up the mock client to expect a SendMessage call with the above input
	// and return the above output
	mockClient.On("SendMessage", &sqs.SendMessageInput{
		MessageBody: aws.String(messageBody),
		QueueUrl:    aws.String(dlqURL),
	}).Return(&sqs.SendMessageOutput{}, nil)

	// Call SendMessageToDLQ and assert it doesn't return an error
	err := p.SendMessageToDLQ(messageBody, dlqURL)

	// Assert that the expectations were met and there's no error
	mockClient.AssertExpectations(t)
	assert.Nil(t, err)
}

// TestDeleteMessage tests the DeleteMessage method of the Poller.
func TestDeleteMessage(t *testing.T) {
	// Create a new mock SQS client and a Poller with this client
	mockClient := new(mockSQSClient)
	p := poller.NewPoller(mockClient, "testQueueURL")

	// Set up the expected input for the DeleteMessage call
	receiptHandle := "testReceiptHandle"

	// Set up the mock client to expect a DeleteMessage call with the above input
	// and return no error
	mockClient.On("DeleteMessage", &sqs.DeleteMessageInput{
		QueueUrl:      aws.String("testQueueURL"),
		ReceiptHandle: aws.String(receiptHandle),
	}).Return(&sqs.DeleteMessageOutput{}, nil)

	// Call DeleteMessage and assert it doesn't return an error
	err := p.DeleteMessage(&receiptHandle)

	// Assert that the expectations were met and there's no error
	mockClient.AssertExpectations(t)
	assert.Nil(t, err)
}
