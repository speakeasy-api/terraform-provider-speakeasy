// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"speakeasy/internal/sdk/pkg/models/shared"
)

type UpsertPluginResponse struct {
	ContentType string
	// Default error response
	Error *shared.Error
	// OK
	Plugin      *shared.Plugin
	StatusCode  int
	RawResponse *http.Response
}
