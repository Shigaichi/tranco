package tranco

import "fmt"

// APIError is an error from Tranco API.
type APIError struct {
	HTTPStatus int
	Code       int
	Message    string
}

func (err *APIError) Error() string {
	if err.Message == "" {
		return fmt.Sprintf("request failed with status code %d", err.HTTPStatus)
	}
	return fmt.Sprintf("StatusCode: %d, Code: %d, Message: %s", err.HTTPStatus, err.Code, err.Message)
}
