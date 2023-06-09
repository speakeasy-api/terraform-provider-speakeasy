// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"speakeasy/internal/sdk/pkg/models/shared"
)

type QueryEventLogRequest struct {
	// The filter to apply to the query.
	Filters *shared.Filters `queryParam:"serialization=json,name=filters"`
}

type QueryEventLogResponse struct {
	// OK
	BoundedRequests []shared.BoundedRequest
	ContentType     string
	// Default error response
	Error       *shared.Error
	StatusCode  int
	RawResponse *http.Response
}
