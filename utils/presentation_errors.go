package utils

import (
	"context"
	"net/http"
	"pvms/pkg/errors"
)

// Errors
var (
	errJSON         = "Invalid JSON Format"
	errUserID       = "Invalid User ID"
	errRole         = "Invalid Role"
	errRequestID    = "Invalid Request ID"
	errInvalidParam = "Invalid Param"
)

type Error struct {
	Status     int    `json:"status"`  // HTTP status code
	StatusText string `json:"code"`    // Internal error code
	Message    string `json:"message"` // Error message for end user
	Internal   string `json:"-"`       // Internal error message (omitted from JSON)
}

func GetPresentationError(err error) Error {
	mErr := errors.Parse(err)
	retCode := http.StatusInternalServerError
	retStatusText := http.StatusText(http.StatusInternalServerError)
	if mErr.GetCode() != 0 {
		retCode = mErr.GetCode()
		retStatusText = mErr.GetStatusText()
	}

	return Error{
		Status:     retCode,
		StatusText: retStatusText, // error codes are not supported yet
		Message:    mErr.GetDetail(),
	}
}

func ErrJSON(ctx context.Context) error {
	return errors.BadRequest(ctx, errJSON)
}

func ErrRequestID(ctx context.Context) error {
	return errors.BadRequest(ctx, errRequestID)
}

func ErrUserID(ctx context.Context) error {
	return errors.BadRequest(ctx, errUserID)
}

func ErrRole(ctx context.Context) error {
	return errors.BadRequest(ctx, errRole)
}

func ErrInvalidParam(ctx context.Context) error {
	return errors.BadRequest(ctx, errInvalidParam)
}
