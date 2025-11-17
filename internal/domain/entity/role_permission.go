package entity

import (
	"github.com/google/uuid"
)

// RolePermission represents the many-to-many relationship between roles and permissions
type RolePermission struct {
	RoleID       uuid.UUID `gorm:"primaryKey"`
	PermissionID uuid.UUID `gorm:"primaryKey"`
}

// TableName specifies the table name for GORM
func (RolePermission) TableName() string {
	return "role_permissions"
}
