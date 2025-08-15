package api_test

import (
	"testing"

	"github.com/axatol/kinde-go/pkg/api"
	"github.com/stretchr/testify/assert"
)

func indirect[T any](t *testing.T, value T) *T {
	t.Helper()
	return &value
}

func TestFormatErrorResponses_Nil(t *testing.T) {
	err := api.FormatErrorResponses(nil)
	assert.NoError(t, err)
}

func TestFormatErrorResponses_UnexpectedType(t *testing.T) {
	err := api.FormatErrorResponses(1)
	assert.ErrorContains(t, err, "request failed with errors: unhandled error (int)")
}

func TestFormatErrorResponses_NotFoundResponseNil(t *testing.T) {
	err := api.FormatErrorResponses((*api.NotFoundResponse)(nil))
	assert.NoError(t, err)
}

func TestFormatErrorResponses_NotFoundResponse(t *testing.T) {
	err := api.FormatErrorResponses(&api.NotFoundResponse{Errors: &struct {
		Code    *string "json:\"code,omitempty\""
		Message *string "json:\"message,omitempty\""
	}{
		Code:    indirect(t, "not_found"),
		Message: indirect(t, "The requested resource was not found"),
	}})
	assert.ErrorContains(t, err, "request failed with errors: [not_found: The requested resource was not found]")
}

func TestFormatErrorResponses_ErrorResponseNil(t *testing.T) {
	err := api.FormatErrorResponses((*api.ErrorResponse)(nil))
	assert.NoError(t, err)
}

func TestFormatErrorResponses_EmptyErrorResponse(t *testing.T) {
	err := api.FormatErrorResponses(&api.ErrorResponse{Errors: &[]api.Error{}})
	assert.ErrorContains(t, err, "request failed with errors: unknown error")
}

func TestFormatErrorResponses_ErrorResponse(t *testing.T) {
	err := api.FormatErrorResponses(&api.ErrorResponse{Errors: &[]api.Error{
		{},
		{
			Code:    indirect(t, "invalid_request"),
			Message: indirect(t, "The request is invalid"),
		},
		{
			Code:    indirect(t, "invalid_request"),
			Message: indirect(t, "The request is invalid"),
		},
	}})
	assert.ErrorContains(t, err, "request failed with errors: [invalid_request: The request is invalid; invalid_request: The request is invalid]")
}
