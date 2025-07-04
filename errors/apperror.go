package errors

import (
	"fmt"
	"net/http"
)

//* struct for custom error
type AppError struct {
	Code int		//* status code 
	Message string	//* error message
	Err	error		//* optional error
}

//* error interface
func (e *AppError) Error() string {
	if e.Err != nil {
        return fmt.Sprintf("%s: %v", e.Message, e.Err)
    }
    return e.Message
}

//* create new custom error
func New(code int, message string, err error) *AppError {
	return &AppError{
		Code: code,
		Message: message,
		Err: err,
	}
}

//* init common custom error
var (
	ErrBadRequest      = New(http.StatusBadRequest, "Bad Request", nil)
    ErrUnauthorized    = New(http.StatusUnauthorized, "Unauthorized", nil)
    ErrNotFound        = New(http.StatusNotFound, "Resource not found", nil)
    ErrInternalServer  = New(http.StatusInternalServerError, "Internal Server Error", nil)
)