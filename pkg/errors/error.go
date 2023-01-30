package errors

import "fmt"

type StatusError struct {
	Code  int
	Cause string
}

func (err StatusError) Error() string {
	return fmt.Sprintf("unexpected status code: %d, cause: %s", err.Code, err.Cause)
}

// ErrUnexpectedStatusCode is returned on unexpected HTTP status codes
type ErrUnexpectedStatusCode int

func (err ErrUnexpectedStatusCode) Error() string {
	return fmt.Sprintf("unexpected status code: %d", err)
}
