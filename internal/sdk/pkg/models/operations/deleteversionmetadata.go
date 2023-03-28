// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"speakeasy/internal/sdk/pkg/models/shared"
)

type DeleteVersionMetadataPathParams struct {
	// The ID of the Api to delete metadata for.
	APIID string `pathParam:"style=simple,explode=false,name=apiID"`
	// The key of the metadata to delete.
	MetaKey string `pathParam:"style=simple,explode=false,name=metaKey"`
	// The value of the metadata to delete.
	MetaValue string `pathParam:"style=simple,explode=false,name=metaValue"`
	// The version ID of the Api to delete metadata for.
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

type DeleteVersionMetadataRequest struct {
	PathParams DeleteVersionMetadataPathParams
}

type DeleteVersionMetadataResponse struct {
	ContentType string
	// Default error response
	Error       *shared.Error
	StatusCode  int
	RawResponse *http.Response
}