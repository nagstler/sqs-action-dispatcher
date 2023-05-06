package actions

import (
	"encoding/json"
)

type Action interface {
	Execute(data json.RawMessage) error
}
