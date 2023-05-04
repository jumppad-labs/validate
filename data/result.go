package data

import (
	"encoding/json"
	"fmt"
	"os"
)

type Result struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message,omitempty"`
	Details map[string]interface{} `json:"details,omitempty"`
}

func Error(err error) {
	result := &Result{
		Status:  254,
		Message: fmt.Sprintf("Error while validating: %s", err.Error()),
		Details: map[string]interface{}{
			"error": err,
		},
	}

	json.NewEncoder(os.Stderr).Encode(result)
	os.Exit(254)
}
