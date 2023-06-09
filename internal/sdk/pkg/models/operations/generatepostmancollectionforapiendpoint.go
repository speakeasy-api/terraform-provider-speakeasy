// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"speakeasy/internal/sdk/pkg/models/shared"
)

type GeneratePostmanCollectionForAPIEndpointRequest struct {
	// The ID of the ApiEndpoint to generate a Postman collection for.
	APIEndpointID string `pathParam:"style=simple,explode=false,name=apiEndpointID"`
	// The ID of the Api to generate a Postman collection for.
	APIID string `pathParam:"style=simple,explode=false,name=apiID"`
	// The version ID of the Api to generate a Postman collection for.
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

type GeneratePostmanCollectionForAPIEndpointResponse struct {
	ContentType string
	// Default error response
	Error *shared.Error
	// OK
	PostmanCollection []byte
	StatusCode        int
	RawResponse       *http.Response
}
