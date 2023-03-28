// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"speakeasy/internal/sdk/pkg/models/shared"
)

type RunPluginPathParams struct {
	// The ID of the plugin to run.
	PluginID string `pathParam:"style=simple,explode=false,name=pluginID"`
}

type RunPluginQueryParams struct {
	// The filter to apply to the query.
	Filters *shared.Filters `queryParam:"serialization=json,name=filters"`
}

type RunPluginRequest struct {
	PathParams  RunPluginPathParams
	QueryParams RunPluginQueryParams
}

type RunPluginResponse struct {
	// OK
	BoundedRequests []shared.BoundedRequest
	ContentType     string
	// Default error response
	Error       *shared.Error
	StatusCode  int
	RawResponse *http.Response
}