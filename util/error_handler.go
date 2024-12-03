package util

import (
	"encoding/json"
	"log"
	"net/http"
)

// APIError represents a standard error response.
type APIError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// WriteError writes an error message to the response.
func WriteError(w http.ResponseWriter, message string, code int) {
	// Set the Content-Type to application/json
	w.Header().Set("Content-Type", "application/json")

	// Set the HTTP status code
	w.WriteHeader(code)

	// Prepare the error response
	apiError := APIError{
		Message: message,
		Code:    code,
	}

	// Encode the error response to JSON and handle any encoding error
	if err := json.NewEncoder(w).Encode(apiError); err != nil {
		// Log the error, as we failed to send the response
		log.Printf("failed to write error response: %v", err)

		// Optionally, send a generic error to the user in case of failure
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
