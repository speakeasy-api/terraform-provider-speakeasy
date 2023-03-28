// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"speakeasy/internal/sdk/pkg/models/shared"
)

type GetRequestFromEventLogRequest struct {
	// The ID of the request to retrieve.
	RequestID string `pathParam:"style=simple,explode=false,name=requestID"`
}

type GetRequestFromEventLogResponse struct {
	ContentType string
	// Default error response
	Error       *shared.Error
	StatusCode  int
	RawResponse *http.Response
	// OK
	UnboundedRequest *shared.UnboundedRequest
}
