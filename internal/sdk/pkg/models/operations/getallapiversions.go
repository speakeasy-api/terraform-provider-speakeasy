// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"speakeasy/internal/sdk/pkg/models/shared"
)

// GetAllAPIVersionsOp - Configuration for filter operations
type GetAllAPIVersionsOp struct {
	// Whether to AND or OR the filters
	And bool `queryParam:"name=and"`
}

type GetAllAPIVersionsRequest struct {
	// The ID of the Api to retrieve.
	APIID string `pathParam:"style=simple,explode=false,name=apiID"`
	// Metadata to filter Apis on
	Metadata map[string][]string `queryParam:"style=deepObject,explode=true,name=metadata"`
	// Configuration for filter operations
	Op *GetAllAPIVersionsOp `queryParam:"style=deepObject,explode=true,name=op"`
}

type GetAllAPIVersionsResponse struct {
	// OK
	Apis        []shared.API
	ContentType string
	// Default error response
	Error       *shared.Error
	StatusCode  int
	RawResponse *http.Response
}
