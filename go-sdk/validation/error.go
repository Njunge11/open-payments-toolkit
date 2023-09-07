package validation

import (
	"encoding/json"
)

type ValidationError struct {
	Name   string `json:"name"`
	Reason string `json:"reason"`
}

type ValidationErrors struct {
	Type          string            `json:"type"`
	Title         string            `json:"title"`
	InvalidParams []ValidationError `json:"invalid-params"`
}

func (ve ValidationErrors) Error() string {
	errJSON, err := json.Marshal(ve)
	if err != nil {
		return err.Error()
	}
	return string(errJSON)
}
