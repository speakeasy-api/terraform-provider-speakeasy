// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"speakeasy/internal/sdk/pkg/models/shared"
)

type GenerateOpenAPISpecForAPIEndpointPathParams struct {
	// The ID of the ApiEndpoint to generate an OpenAPI specification for.
	APIEndpointID string `pathParam:"style=simple,explode=false,name=apiEndpointID"`
	// The ID of the Api to generate an OpenAPI specification for.
	APIID string `pathParam:"style=simple,explode=false,name=apiID"`
	// The version ID of the Api to generate an OpenAPI specification for.
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

type GenerateOpenAPISpecForAPIEndpointRequest struct {
	PathParams GenerateOpenAPISpecForAPIEndpointPathParams
}

type GenerateOpenAPISpecForAPIEndpointResponse struct {
	ContentType string
	// Default error response
	Error *shared.Error
	// OK
	GenerateOpenAPISpecDiff *shared.GenerateOpenAPISpecDiff
	StatusCode              int
	RawResponse             *http.Response
}
