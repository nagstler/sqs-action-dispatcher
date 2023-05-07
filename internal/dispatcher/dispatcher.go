package dispatcher

import (
	"encoding/json"
	"fmt"

	"github.com/nagstler/sqs-action-dispatcher/internal/actions"
)

// DispatchMessage represents the structure of an SQS message used for dispatching actions.
type DispatchMessage struct {
	Action string          `json:"action"`
	Data   json.RawMessage `json:"data"`
}

// Dispatcher is responsible for dispatching actions based on the received messages.
type Dispatcher struct{}

// NewDispatcher creates a new Dispatcher instance.
func NewDispatcher() *Dispatcher {
	return &Dispatcher{}
}

// Dispatch processes the given message string, extracts the action and data,
// and executes the appropriate action.
func (d *Dispatcher) Dispatch(msg string) error {
	var dispatchMessage DispatchMessage
	err := json.Unmarshal([]byte(msg), &dispatchMessage)
	if err != nil {
		return fmt.Errorf("failed to unmarshal dispatch message: %w", err)
	}

	var action actions.Action

	// Determine the action to be executed based on the action specified in the message
	switch dispatchMessage.Action {
	case "sns_publish":
		action = &actions.SNSAction{}
	default:
		return fmt.Errorf("unknown action: %s", dispatchMessage.Action)
	}

	// Execute the action with the provided data
	err = action.Execute(dispatchMessage.Data)
	if err != nil {
		return fmt.Errorf("failed to execute action: %w", err)
	}

	return nil
}
