package entity

import (
	"time"

	"github.com/google/uuid"
)

// Permission represents a permission entity in the domain
type Permission struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`      // e.g., "user:read", "user:update", "dorm:read"
	Slug      string    `json:"slug"`
	Resource  string    `json:"resource"`  // e.g., "user", "dorm"
	Action    string    `json:"action"`     // e.g., "read", "update", "create", "delete"
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relations
	Roles []Role `gorm:"many2many:role_permissions;" json:"roles,omitempty"`
}

// TableName specifies the table name for GORM
func (Permission) TableName() string {
	return "permissions"
}
