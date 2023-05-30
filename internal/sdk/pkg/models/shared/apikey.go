// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// APIKey - A respresentation of an ApiKey.
type APIKey struct {
	// The id of the ApiKey. Use this to delete it.
	ID string `json:"id"`
	// The key of the ApiKey.
	Key string `json:"key"`
	// The name of the ApiKey.
	Name string `json:"name"`
	// The workspace id of the ApiKey.
	WorkspaceID string `json:"workspace_id"`
}