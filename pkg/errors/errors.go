// Package errors - is the place where our implementation of error methods/messages lives
package errors

import (
	"net/http"
)

const (
	// ErrorDefault represents an unspecified error.
	ErrorDefault = Error(iota)
	// ErrorMarshalData represents an error while marshaling data.
	ErrorMarshalData
	// ErrorUnmarshalData represents an error while unmarshaling data.
	ErrorUnmarshalData
	// ErrorNoData occurs when no information is returned a service.
	ErrorNoData
	// ErrorInsert represents an error when inserting a record.
	ErrorInsert
	// ErrorUpdate represents an error when updating a record.
	ErrorUpdate
	// ErrorDelete represents an error when deleting a record.
	ErrorDelete
	// ErrorInvalidInput occurs when the input is invalid.
	ErrorInvalidInput
	// ErrorUnprocessableEntity occurs when the request is not processable.
	ErrorUnprocessableEntity
	// ErrorUnauthorized occurs the user lacks valid authentication credentials for the requested resource.
	ErrorUnauthorized
	// ErrorForbidden occurs the user does not have access to the requested resource.
	ErrorForbidden
	// ErrorBadRequest occurs when the request could not be understood by the server due to malformed syntax.
	ErrorBadRequest
	// ErrorMethodNotAllowed occurs when the request uses a method that is not allowed.
	ErrorMethodNotAllowed
	// ErrorReferentialIntegrity occurs when there is a foreign key constraint.
	ErrorReferentialIntegrity
	// ErrorRateLimit occurs when the user has sent too many requests in a given amount of time.
	ErrorRateLimit
	// ErrorDuplicateConstraintViolation occurs on an attempt to insert a row that violates a unique constraint.
	ErrorDuplicateConstraintViolation
	// ErrorPreconditionFailed occurs when one or more conditions given in the request evaluated to false.
	ErrorPreconditionFailed
	// ErrorNotAcceptable indicates that the server cannot produce a response matching the list of acceptable values
	// defined in the request's proactive content negotiation headers, and that the server is unwilling to supply a default representation
	ErrorNotAcceptable
)

// Error represents an error encountered using a service.
type Error int

// Error returns the error message associated with the Error.
// nolint
func (e Error) Error() string {
	switch e {
	case ErrorReferentialIntegrity:
		return "error referential integrity"
	case ErrorMarshalData:
		return "error while marshaling data"
	case ErrorUnmarshalData:
		return "error while unmarshaling data"
	case ErrorBadRequest:
		return http.StatusText(http.StatusBadRequest)
	case ErrorNoData:
		return http.StatusText(http.StatusNotFound)
	case ErrorInsert:
		return "error while inserting this record"
	case ErrorUpdate:
		return "error while updating this record"
	case ErrorDelete:
		return "error while deleting this record"
	case ErrorInvalidInput:
		return http.StatusText(http.StatusBadRequest)
	case ErrorUnprocessableEntity:
		return http.StatusText(http.StatusUnprocessableEntity)
	case ErrorUnauthorized:
		return http.StatusText(http.StatusUnauthorized)
	case ErrorForbidden:
		return http.StatusText(http.StatusForbidden)
	case ErrorMethodNotAllowed:
		return http.StatusText(http.StatusMethodNotAllowed)
	case ErrorRateLimit:
		return http.StatusText(http.StatusTooManyRequests)
	case ErrorDuplicateConstraintViolation:
		return "error duplicate constraint violation"
	case ErrorPreconditionFailed:
		return http.StatusText(http.StatusPreconditionFailed)
	case ErrorNotAcceptable:
		return http.StatusText(http.StatusNotAcceptable)
	default:
		return http.StatusText(http.StatusInternalServerError)
	}
}

// StatusCode returns the default http status code for the error.
// nolint
func (e Error) StatusCode() int {
	switch e {
	case ErrorBadRequest:
		return http.StatusBadRequest
	case ErrorNoData:
		return http.StatusNotFound
	case ErrorInvalidInput:
		return http.StatusBadRequest
	case ErrorUnprocessableEntity:
		return http.StatusUnprocessableEntity
	case ErrorUnauthorized:
		return http.StatusUnauthorized
	case ErrorForbidden:
		return http.StatusForbidden
	case ErrorMethodNotAllowed:
		return http.StatusMethodNotAllowed
	case ErrorRateLimit:
		return http.StatusTooManyRequests
	case ErrorDuplicateConstraintViolation:
		return http.StatusConflict
	case ErrorPreconditionFailed:
		return http.StatusPreconditionFailed
	case ErrorNotAcceptable:
		return http.StatusNotAcceptable
	default:
		return http.StatusInternalServerError
	}
}
