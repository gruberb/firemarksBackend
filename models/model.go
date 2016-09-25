package models

type (
	// ValidationError represents a failed validation
	ValidationError struct {
		Message string `json:"message"`
	}
)
