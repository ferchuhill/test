package client

import "fmt"

// APIError represents an error that can occurred while calling the API.
type APIError struct {
	// Error class
	Class string `json:"class,omitempty"`
	// Error message.
	Message string `json:"message"`
	// Error details
	Details map[string]string `json:"details,omitempty"`
	// HTTP code.
	Code int
}

// Print the error with a specific formats
func (err *APIError) Error() string {
	return fmt.Sprintf("Error %d: %q", err.Code, err.Message)
}
