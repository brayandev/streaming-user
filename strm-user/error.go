package user

import "fmt"

// ErrorType error types
type ErrorType int

// ErrorType error type values.
const (
	ErrorUnknown ErrorType = iota
	ErrorInvalidContent
	ErrorUnprocessableEntity
)

// ErrorCodes
const (
	ErrorCodeUnknown string = "STRU0000"

	// Common errors STRUxxx.
	ErrorCodeInvalidContent      = "STRU1001"
	ErrorCodeUnprocessableEntity = "STRU1002"
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

// NewInvalidContentError constructor of strm-user invalid content error
func NewInvalidContentError(message string) *Error {
	return NewError(ErrorCodeInvalidContent, message, ErrorInvalidContent)
}

// NewUnprocessableEntityError constructor of strm-user unprocessable entity error.
func NewUnprocessableEntityError(message string) *Error {
	return NewError(ErrorCodeUnprocessableEntity, message, ErrorUnprocessableEntity)
}

// Error return a string representation of and Error.
func (e Error) Error() string {
	return fmt.Sprintf("%s - %s", e.ErrCode, e.Message)
}

// Version represents a version of error.
func (e Error) Version() string {
	return "strm-user.error.v1"
}
