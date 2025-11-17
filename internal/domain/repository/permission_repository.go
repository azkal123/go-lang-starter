package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/your-org/go-backend-starter/internal/domain/entity"
)

// PermissionRepository defines the interface for permission data operations
type PermissionRepository interface {
	Create(ctx context.Context, permission *entity.Permission) error
	GetByID(ctx context.Context, id uuid.UUID) (*entity.Permission, error)
	GetBySlug(ctx context.Context, slug string) (*entity.Permission, error)
	Update(ctx context.Context, permission *entity.Permission) error
	Delete(ctx context.Context, id uuid.UUID) error
	List(ctx context.Context, limit, offset int) ([]*entity.Permission, int64, error)
}
