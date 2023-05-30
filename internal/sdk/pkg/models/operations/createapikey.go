// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"speakeasy/internal/sdk/pkg/models/shared"
)

type CreateAPIKeyRequest struct {
	APIKeyCreate *shared.APIKeyCreate `request:"mediaType=application/json"`
	// The ID of the workspace to create the ApiKey for.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspaceID"`
}

type CreateAPIKeyResponse struct {
	// OK
	APIKey      *shared.APIKey
	ContentType string
	// Default error response
	Error       *shared.Error
	StatusCode  int
	RawResponse *http.Response
}