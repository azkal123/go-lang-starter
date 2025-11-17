package service

import (
	"context"

	"github.com/google/uuid"
)

// AuthService defines the interface for authentication operations
type AuthService interface {
	AssignRoleToUser(ctx context.Context, userID, roleID uuid.UUID) error
	RemoveRoleFromUser(ctx context.Context, userID, roleID uuid.UUID) error
	GetUserRoles(ctx context.Context, userID uuid.UUID) ([]string, error)
	GetUserPermissions(ctx context.Context, userID uuid.UUID) ([]string, error)
}
