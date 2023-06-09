// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

// EmbedToken - A representation of an embed token granted for working with Speakeasy components.
type EmbedToken struct {
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at"`
	// The ID of the user that created this token.
	CreatedBy string `json:"created_by"`
	// A detailed description of the EmbedToken.
	Description string `json:"description"`
	// The time this token expires.
	ExpiresAt time.Time `json:"expires_at"`
	// The filters applied to this token.
	Filters string `json:"filters"`
	// The ID of this EmbedToken.
	ID string `json:"id"`
	// The last time this token was used.
	LastUsed *time.Time `json:"last_used,omitempty"`
	// The time this token was revoked.
	RevokedAt *time.Time `json:"revoked_at,omitempty"`
	// The ID of the user that revoked this token.
	RevokedBy *string `json:"revoked_by,omitempty"`
	// The workspace ID this token belongs to.
	WorkspaceID string `json:"workspace_id"`
}
