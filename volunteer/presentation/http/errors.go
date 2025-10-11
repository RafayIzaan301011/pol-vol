package http

import (
	"context"
	"pvms/pkg/errors"
)

// Errors
const (
	errJSON         = "Invalid JSON Format"
	errUserID       = "Invalid User ID"
	errRole         = "Invalid Role"
	errRequestID    = "Invalid Request ID"
	errInvalidParam = "Invalid Param"
)

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
