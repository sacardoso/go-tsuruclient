/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package tsuru

import (
	"time"
)

// An authorization token associated to a team.
type TeamToken struct {
	Token string `json:"token,omitempty"`

	TokenId string `json:"token_id,omitempty"`

	Description string `json:"description,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`

	ExpiresAt time.Time `json:"expires_at,omitempty"`

	LastAccess time.Time `json:"last_access,omitempty"`

	CreatorEmail string `json:"creator_email,omitempty"`

	Team string `json:"team,omitempty"`

	Roles []RoleInstance `json:"roles,omitempty"`
}
