package merche

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// ExVeError struct for ExVeError.
type ExVeError struct {
	ExveErrorID  *string `json:"exveErrorId,omitempty"`
	ExveErrorMsg *string `json:"exveErrorMsg,omitempty"`
	ExveErrorRef *string `json:"exveErrorRef,omitempty"`
}

func (e *ExVeError) Error() string {
	return ""
}

type UnauthorizedError struct {
	ErrorMessage *string `json:"errorMessage,omitempty"`
	StatusCode   *string `json:"statusCode,omitempty"`
	Message      *string `json:"message,omitempty"`
}

func (e *UnauthorizedError) Error() string {
	return ""
}

type MercedesAPIError struct {
	StatusCode int
}

func (e *MercedesAPIError) Error() string {
	return http.StatusText(e.StatusCode)
}

func handleResponseError(resp *http.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("error handler: error reading response body")
	}

	if isExVeError(resp.StatusCode) {
		var exVeError ExVeError
		err = json.Unmarshal(body, &exVeError)
		if err != nil {
			return errors.New("error handler: error unmarshalling response body")
		}
		return &exVeError
	}

	if resp.StatusCode == http.StatusUnauthorized {
		var authErr UnauthorizedError
		err = json.Unmarshal(body, &authErr)
		if err != nil {
			return errors.New("error handler: error unmarshalling response body")
		}
		return &authErr
	}

	return &MercedesAPIError{
		resp.StatusCode,
	}
}

func isExVeError(statusCode int) bool {
	return statusCode == http.StatusBadRequest || statusCode == http.StatusForbidden ||
		statusCode == http.StatusInternalServerError || statusCode == http.StatusServiceUnavailable
}
