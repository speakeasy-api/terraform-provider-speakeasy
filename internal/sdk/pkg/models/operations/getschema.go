// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"speakeasy/internal/sdk/pkg/models/shared"
)

type GetSchemaRequest struct {
	// The ID of the Api to get the schema for.
	APIID string `pathParam:"style=simple,explode=false,name=apiID"`
	// The version ID of the Api to delete metadata for.
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

type GetSchemaResponse struct {
	ContentType string
	// Default error response
	Error *shared.Error
	// OK
	Schema      *shared.Schema
	StatusCode  int
	RawResponse *http.Response
}
