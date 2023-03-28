// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

// UnboundedRequest - An UnboundedRequest represents the HAR content capture by Speakeasy when logging a request.
type UnboundedRequest struct {
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at"`
	// The HAR content of the request.
	Har string `json:"har"`
	// The size of the HAR content in bytes.
	HarSizeBytes int64 `json:"har_size_bytes"`
	// The ID of this request.
	RequestID string `json:"request_id"`
	// The workspace ID this request was made to.
	WorkspaceID string `json:"workspace_id"`
}