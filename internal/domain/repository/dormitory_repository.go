package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/your-org/go-backend-starter/internal/domain/entity"
)

// DormitoryRepository defines the interface for dormitory data operations
type DormitoryRepository interface {
	Create(ctx context.Context, dormitory *entity.Dormitory) error
	GetByID(ctx context.Context, id uuid.UUID) (*entity.Dormitory, error)
	Update(ctx context.Context, dormitory *entity.Dormitory) error
	Delete(ctx context.Context, id uuid.UUID) error
	List(ctx context.Context, limit, offset int) ([]*entity.Dormitory, int64, error)
	AssignToUser(ctx context.Context, userID, dormitoryID uuid.UUID) error
	RemoveFromUser(ctx context.Context, userID, dormitoryID uuid.UUID) error
	GetUserDormitories(ctx context.Context, userID uuid.UUID) ([]*entity.Dormitory, error)
}
