// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"speakeasy/internal/sdk/pkg/models/shared"
)

type DeleteSchemaRequest struct {
	// The ID of the Api to delete schemas for.
	APIID string `pathParam:"style=simple,explode=false,name=apiID"`
	// The revision ID of the schema to delete.
	RevisionID string `pathParam:"style=simple,explode=false,name=revisionID"`
	// The version ID of the Api to delete metadata for.
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

type DeleteSchemaResponse struct {
	ContentType string
	// Default error response
	Error       *shared.Error
	StatusCode  int
	RawResponse *http.Response
}
