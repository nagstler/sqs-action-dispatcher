package dispatcher

import (
	"encoding/json"
	"fmt"

	"github.com/nagstler/sqs-action-dispatcher/internal/actions"
)

type DispatchMessage struct {
	Action string          `json:"action"`
	Data   json.RawMessage `json:"data"`
}

type Dispatcher struct{}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{}
}

func (d *Dispatcher) Dispatch(msg string) error {
	var dispatchMessage DispatchMessage
	err := json.Unmarshal([]byte(msg), &dispatchMessage)
	if err != nil {
		return fmt.Errorf("failed to unmarshal dispatch message: %w", err)
	}

	var action actions.Action

	switch dispatchMessage.Action {
	case "sns_publish":
		action = &actions.SNSAction{}
	default:
		return fmt.Errorf("unknown action: %s", dispatchMessage.Action)
	}

	err = action.Execute(dispatchMessage.Data)
	if err != nil {
		return fmt.Errorf("failed to execute action: %w", err)
	}

	return nil
}
