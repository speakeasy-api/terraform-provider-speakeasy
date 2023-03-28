// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

// Schema - A Schema represents an API schema for a particular Api and Version.
type Schema struct {
	// The ID of the Api this Schema belongs to.
	APIID string `json:"api_id"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at"`
	// A detailed description of the Schema.
	Description string `json:"description"`
	// An ID referencing this particular revision of the Schema.
	RevisionID string `json:"revision_id"`
	// The version ID of the Api this Schema belongs to.
	VersionID string `json:"version_id"`
	// The workspace ID this Schema belongs to.
	WorkspaceID string `json:"workspace_id"`
}