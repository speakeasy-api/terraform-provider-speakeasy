// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

// VersionMetadataInput - A set of keys and associated values, attached to a particular version of an Api.
type VersionMetadataInput struct {
	// The key for this metadata.
	MetaKey string `json:"meta_key"`
	// One of the values for this metadata.
	MetaValue string `json:"meta_value"`
}

// VersionMetadata - A set of keys and associated values, attached to a particular version of an Api.
type VersionMetadata struct {
	// The ID of the Api this Metadata belongs to.
	APIID string `json:"api_id"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at"`
	// The key for this metadata.
	MetaKey string `json:"meta_key"`
	// One of the values for this metadata.
	MetaValue string `json:"meta_value"`
	// The version ID of the Api this Metadata belongs to.
	VersionID string `json:"version_id"`
	// The workspace ID this Metadata belongs to.
	WorkspaceID string `json:"workspace_id"`
}