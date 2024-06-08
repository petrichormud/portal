package handler

import "net/http"

type Error struct {
	Status int
}

func (e *Error) Error() string {
	return "handler error"
}

var (
	ErrNotFound     = &Error{Status: http.StatusNotFound}
	ErrInternal     = &Error{Status: http.StatusInternalServerError}
	ErrUnauthorized = &Error{Status: http.StatusUnauthorized}
)
