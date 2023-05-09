# SQS Action Dispatcher
[![License](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT) [![CI Build](https://github.com/nagstler/sqs-action-dispatcher/actions/workflows/main.yml/badge.svg?branch=main)](https://github.com/nagstler/sqs-action-dispatcher/actions/workflows/main.yml) [![Maintainability](https://api.codeclimate.com/v1/badges/fe760284be051623a2d4/maintainability)](https://codeclimate.com/github/nagstler/sqs-action-dispatcher/maintainability) [![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](http://makeapullrequest.com)

The SQS Action Dispatcher is a Go application that polls messages from an AWS Simple Queue Service (SQS) queue, processes the messages concurrently using worker goroutines, and dispatches actions based on the message content.

## Features
- Polls messages from an SQS queue in batches for improved throughput
- Processes messages concurrently using worker goroutines
- Moves failed messages to a Dead Letter Queue (DLQ) for further inspection
- Dispatches actions based on message content (e.g., send an SNS notification)

## Prerequisites
- Go 1.16 or higher
- AWS account with an SQS queue and optional DLQ configured
- AWS CLI or environment variables with access to the SQS queue

## Getting Started
1. Clone the repository:

```sh

git clone https://github.com/your_username/sqs-action-dispatcher.git
```


1. Change directory to the project:

```sh

cd sqs-action-dispatcher
```


1. Build the project:

```sh

go build
```


1. Set the AWS credentials and region as environment variables:

```sh

export AWS_ACCESS_KEY_ID=your_access_key
export AWS_SECRET_ACCESS_KEY=your_secret_key
export AWS_REGION=your_aws_region
```


1. Set environment variables for the AWS SQS queue URL and DLQ URL:

```sh

export SQS_QUEUE_URL=https://sqs.your_aws_region.amazonaws.com/your_account_id/your_queue_name
export SQS_DLQ_URL=https://sqs.your_aws_region.amazonaws.com/your_account_id/your_dlq_name
```


1. Run the application:

```sh

./sqs-action-dispatcher
```



The application will start polling messages from the configured SQS queue, dispatch actions based on the message content, and move failed messages to the DLQ.
## SNS Action

The SNS action sends a message to an SNS topic. To use the SNS action, your messages should have the following format:

```json

{
  "type": "sns",
  "data": {
    "topic_arn": "arn:aws:sns:your_aws_region:your_account_id:your_topic_name",
    "message": "Your message to send"
  }
}
```

When the SQS Action Dispatcher receives a message with this format, it will send the specified message to the SNS topic.

## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/nagstler/sqs-action-dispatcher. This project is intended to be a safe, welcoming space for collaboration, and contributors are expected to adhere to the [code of conduct](https://github.com/nagstler/sqs-action-dispatcher/blob/main/CODE_OF_CONDUCT.md).

## License

The gem is available as open source under the terms of the [MIT License](https://opensource.org/licenses/MIT).

## Code of Conduct

Everyone interacting in the Chatgpt::Ruby project's codebases, issue trackers, chat rooms and mailing lists is expected to follow the [code of conduct](https://github.com/nagstler/chatgpt-ruby/blob/main/CODE_OF_CONDUCT.md).