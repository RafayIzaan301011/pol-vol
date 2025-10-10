package errors

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"google.golang.org/grpc/status"
)

// MicroError is our interface for micro errors.
type MicroError interface {
	GetID() string
	GetCode() int
	GetDetail() string
	GetStatusText() string
	GetAdditionalInfo() map[string]string
	Error() string
}

// MError implements the micro error interface.
type MError struct {
	ID             string            `json:"id"`
	Code           int               `json:"code"`
	Detail         string            `json:"detail"`
	Status         string            `json:"status"`
	AdditionalInfo map[string]string `json:"additional_info,omitempty"`
}

// Error will return the marshaled error struct.
func (e *MError) Error() string {
	b, err := json.Marshal(e)
	if err != nil {
		return e.Detail
	}

	return string(b)
}

// GetID will return the error id.
func (e *MError) GetID() string {
	return e.ID
}

// GetDetail will return the error detail.
func (e *MError) GetDetail() string {
	return e.Detail
}

// GetCode will return the error code.
func (e *MError) GetCode() int {
	return e.Code
}

// GetStatusText will return the error status text.
func (e *MError) GetStatusText() string {
	return http.StatusText(e.Code)
}

// GetAdditionalInfo will return additional error info.
func (e *MError) GetAdditionalInfo() map[string]string {
	return e.AdditionalInfo
}

// AdditionalInfo is additional information to return for an error.
type AdditionalInfo struct {
	Key   string
	Value string
}

// New create a new MicroError.
func New(ctx context.Context, detail string, code int, additionalInfos ...AdditionalInfo) error {
	return newMError(ctx, code, detail, http.StatusText(code), additionalInfos)
}

// nolint
func newMError(ctx context.Context, code int, detail, status string, additionalInfos []AdditionalInfo) *MError {
	e := &MError{
		Code:           code,
		Detail:         detail,
		Status:         status,
		AdditionalInfo: map[string]string{},
	}
	for _, additionalInfo := range additionalInfos {
		e.AdditionalInfo[additionalInfo.Key] = additionalInfo.Value
	}
	return e
}

// BadRequest (400) The request could not be fulfilled due to the incorrect syntax of the request.
func BadRequest(ctx context.Context, detail string, additionalInfos ...AdditionalInfo) error {
	return newMError(ctx, http.StatusBadRequest,
		detail, http.StatusText(http.StatusBadRequest), additionalInfos)
}

// Unauthorized (401) The requester is not authorized to access the resource. This is similar to 403 but is
// used in cases where authentication is expected but has failed or has not been provided.
func Unauthorized(ctx context.Context, detail string, additionalInfos ...AdditionalInfo) error {
	return newMError(ctx, http.StatusUnauthorized,
		detail, http.StatusText(http.StatusUnauthorized), additionalInfos)
}

// Forbidden (403) The request was formatted correctly but the server is refusing to supply the requested
// resource. Unlike 401, authenticating will not make a difference in the server's response.
func Forbidden(ctx context.Context, detail string, additionalInfos ...AdditionalInfo) error {
	return newMError(ctx, http.StatusForbidden,
		detail, http.StatusText(http.StatusForbidden), additionalInfos)
}

// NotFound (404) The resource could not be found. This is often used as a catch-all for all invalid URIs
// requested of the server.
func NotFound(ctx context.Context, detail string, additionalInfos ...AdditionalInfo) error {
	return newMError(ctx, http.StatusNotFound,
		detail, http.StatusText(http.StatusNotFound), additionalInfos)
}

// NotFoundf (404) The resource could not be found. This is often used as a catch-all for all invalid URIs
// requested of the server.
func NotFoundf(ctx context.Context, detailFormat string, a ...any) error {
	return newMError(ctx, http.StatusNotFound,
		fmt.Sprintf(detailFormat, a...), http.StatusText(http.StatusNotFound), []AdditionalInfo{})
}

// MethodNotAllowed (405) The request Method is not supported by the endpoint.
func MethodNotAllowed(ctx context.Context, detail string, additionalInfos ...AdditionalInfo) error {
	return newMError(ctx, http.StatusMethodNotAllowed,
		detail, http.StatusText(http.StatusMethodNotAllowed), additionalInfos)
}

// NotAcceptable (406) The resource is valid, but cannot be provided in a format specified in the Accept
// headers in the request.
func NotAcceptable(ctx context.Context, detail string, additionalInfos ...AdditionalInfo) error {
	return newMError(ctx, http.StatusNotAcceptable,
		detail, http.StatusText(http.StatusNotAcceptable), additionalInfos)
}

// Conflict (409) The request cannot be completed due to a conflict in the request parameters.
func Conflict(ctx context.Context, detail string, additionalInfos ...AdditionalInfo) error {
	return newMError(ctx, http.StatusConflict,
		detail, http.StatusText(http.StatusConflict), additionalInfos)
}

// Gone (410) The request cannot be completed due to an expired resource.
func Gone(ctx context.Context, detail string, additionalInfos ...AdditionalInfo) error {
	return newMError(ctx, http.StatusGone,
		detail, http.StatusText(http.StatusGone), additionalInfos)
}

// PreconditionFailed (412) The request cannot be completed due
// one or more conditions given in the request evaluated to false.
func PreconditionFailed(ctx context.Context, detail string, additionalInfos ...AdditionalInfo) error {
	return newMError(ctx, http.StatusPreconditionFailed,
		detail, http.StatusText(http.StatusPreconditionFailed), additionalInfos)
}

// UnsupportedMediaType (415) The client provided data with a media type that the server does not support.
func UnsupportedMediaType(ctx context.Context, detail string, additionalInfos ...AdditionalInfo) error {
	return newMError(ctx, http.StatusUnsupportedMediaType,
		detail, http.StatusText(http.StatusUnsupportedMediaType), additionalInfos)
}

// UnprocessableEntity (422) The requester is not authorized to access the resource. This is similar to 403
// but is used in cases where authentication is expected but has failed or has not been provided.
func UnprocessableEntity(ctx context.Context, detail string, additionalInfos ...AdditionalInfo) error {
	return newMError(ctx, http.StatusUnprocessableEntity,
		detail, http.StatusText(http.StatusUnprocessableEntity), additionalInfos)
}

// TooManyRequests (429) The user has sent too many requests in a given amount of time ("rate limiting").
func TooManyRequests(ctx context.Context, detail string, additionalInfos ...AdditionalInfo) error {
	return newMError(ctx, http.StatusTooManyRequests,
		detail, http.StatusText(http.StatusTooManyRequests), additionalInfos)
}

// UnavailableForLegalReasons (451) A server operator has received a legal demand to deny access to a resource
// or to a set of resources that includes the requested resource.
func UnavailableForLegalReasons(ctx context.Context, detail string, additionalInfos ...AdditionalInfo) error {
	return newMError(ctx, http.StatusUnavailableForLegalReasons,
		detail, http.StatusText(http.StatusUnavailableForLegalReasons), additionalInfos)
}

// RequestTimedOut (408) A server operator has decided to close the connection rather than continue waiting.
func RequestTimedOut(ctx context.Context, detail string, additionalInfos ...AdditionalInfo) error {
	return newMError(ctx, http.StatusRequestTimeout,
		detail, http.StatusText(http.StatusRequestTimeout), additionalInfos)
}

// InternalServerError (500) A generic status for an error in the server itself.
func InternalServerError(ctx context.Context, detail string, additionalInfos ...AdditionalInfo) error {
	return newMError(ctx, http.StatusInternalServerError,
		detail, http.StatusText(http.StatusInternalServerError), additionalInfos)
}

// InternalServerErrorf (500) A generic status for an error in the server itself.
func InternalServerErrorf(ctx context.Context, detailFormat string, a ...any) error {
	return newMError(ctx, http.StatusInternalServerError,
		fmt.Sprintf(detailFormat, a...), http.StatusText(http.StatusInternalServerError), []AdditionalInfo{})
}

// NewPlain creates a new error with a plain string.
func NewPlain(str string) error {
	return errors.New(str)
}

// Parse a MicroError from a error.
func Parse(err error) MicroError {
	msg := status.Convert(err).Message()
	e := new(MError)
	handleErr := json.Unmarshal([]byte(msg), e)
	if handleErr != nil {
		e.Detail = msg
	}
	return e
}
