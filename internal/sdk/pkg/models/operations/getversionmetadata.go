// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"speakeasy/internal/sdk/pkg/models/shared"
)

type GetVersionMetadataRequest struct {
	// The ID of the Api to retrieve metadata for.
	APIID string `pathParam:"style=simple,explode=false,name=apiID"`
	// The version ID of the Api to retrieve metadata for.
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

type GetVersionMetadataResponse struct {
	ContentType string
	// Default error response
	Error       *shared.Error
	StatusCode  int
	RawResponse *http.Response
	// OK
	VersionMetadata []shared.VersionMetadata
}
