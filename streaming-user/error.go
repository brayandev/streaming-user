package user

import "fmt"

// ErrorType error types
type ErrorType int

// ErrorType error type values.
const (
	ErrorUnknown ErrorType = iota
	ErrorInvalidContent
)

// ErrorCodes
const (
	ErrorCodeUnknown string = "STRM0000"

	// Common errors STRMxxx.
	ErrorCodeInvalidContent = "STRM1001"
)

// Error representation of error.
type Error struct {
	ErrCode string    `json:"error,omitempty"`
	Message string    `json:"message,omitempty"`
	ErrType ErrorType `json:"-"`
}

// NewError is a constructor of error.
func NewError(code, message string, errType ErrorType) *Error {
	return &Error{ErrCode: code, Message: message, ErrType: errType}
}

// NewUnknownError constructor of permission unknown error.
func NewUnknownError(message string) *Error {
	return NewError(ErrorCodeUnknown, message, ErrorUnknown)
}

// NewInvalidContentError constructor of jobad-inspector invalid content error
func NewInvalidContentError(message string) *Error {
	return NewError(ErrorCodeInvalidContent, message, ErrorInvalidContent)
}

// Error return a string representation of and Error.
func (e Error) Error() string {
	return fmt.Sprintf("%s - %s", e.ErrCode, e.Message)
}

// Version represents a version of error.
func (e Error) Version() string {
	return "strm-user.error.v1"
}
