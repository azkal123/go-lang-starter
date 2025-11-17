package entity

import (
	"github.com/google/uuid"
)

// UserDormitory represents the many-to-many relationship between users and dormitories
// This is the guard mechanism - users can access specific dormitories
type UserDormitory struct {
	UserID      uuid.UUID `gorm:"primaryKey"`
	DormitoryID uuid.UUID `gorm:"primaryKey"`
}

// TableName specifies the table name for GORM
func (UserDormitory) TableName() string {
	return "user_dormitories"
}
