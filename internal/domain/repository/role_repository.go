package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/your-org/go-backend-starter/internal/domain/entity"
)

// RoleRepository defines the interface for role data operations
type RoleRepository interface {
	Create(ctx context.Context, role *entity.Role) error
	GetByID(ctx context.Context, id uuid.UUID) (*entity.Role, error)
	GetBySlug(ctx context.Context, slug string) (*entity.Role, error)
	Update(ctx context.Context, role *entity.Role) error
	Delete(ctx context.Context, id uuid.UUID) error
	List(ctx context.Context, limit, offset int) ([]*entity.Role, int64, error)
	GetWithPermissions(ctx context.Context, id uuid.UUID) (*entity.Role, error)
	AssignPermission(ctx context.Context, roleID, permissionID uuid.UUID) error
	RemovePermission(ctx context.Context, roleID, permissionID uuid.UUID) error
}
