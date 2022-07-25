package merche

import (
	"fmt"
	"net/http"
)

// ExVeError struct for ExVeError.
type ExVeError struct {
	ExveErrorID  string `json:"exveErrorId,omitempty"`
	ExveErrorMsg string `json:"exveErrorMsg,omitempty"`
	ExveErrorRef string `json:"exveErrorRef,omitempty"`
}

func (e *ExVeError) Error() string {
	return fmt.Sprintf("Mercedes API response with %v: %v", e.ExveErrorID, e.ExveErrorMsg)
}

type UnauthorizedError struct {
	ErrorMessage string `json:"errorMessage,omitempty"`
	StatusCode   string `json:"statusCode,omitempty"`
	Message      string `json:"message,omitempty"`
}

func (e *UnauthorizedError) Error() string {
	return fmt.Sprintf("Mercedes API response with %v: %v", e.StatusCode, e.Message)
}

type MercedesAPIError struct {
	StatusCode int
}

func (e *MercedesAPIError) Error() string {
	return http.StatusText(e.StatusCode)
}

func isExVeError(statusCode int) bool {
	return statusCode == http.StatusBadRequest || statusCode == http.StatusForbidden ||
		statusCode == http.StatusInternalServerError || statusCode == http.StatusServiceUnavailable
}
