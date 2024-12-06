package apperror

import "net/http"

var (
	ErrInternalServerError = newError(http.StatusInternalServerError, "Internal Server Error")
	ErrNotFound            = newError(http.StatusNotFound, "Not Found")
	ErrBadRequest          = newError(http.StatusBadRequest, "Bad Request")
)

type Error struct {
	Code    int
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

func newError(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}
