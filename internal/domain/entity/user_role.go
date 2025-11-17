package entity

import (
	"github.com/google/uuid"
)

// UserRole represents the many-to-many relationship between users and roles
type UserRole struct {
	UserID uuid.UUID `gorm:"primaryKey"`
	RoleID uuid.UUID `gorm:"primaryKey"`
}

// TableName specifies the table name for GORM
func (UserRole) TableName() string {
	return "user_roles"
}
