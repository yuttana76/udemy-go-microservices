package errs

import "net/http"

type AppError struct {
	Code    int    `json:",omitempty"`
	Message string `json:"message"`
}

func (e AppError) Error() string {
	return e.Message
}

func NewNotFoundError(message string) error {
	return AppError{Code: http.StatusNotFound, Message: message}
}

func NewInternalServerError() error {
	return AppError{Code: http.StatusInternalServerError, Message: "internal server error"}
}
func NewValidationError(message string) error {
	return AppError{Code: http.StatusUnprocessableEntity, Message: message}
}
