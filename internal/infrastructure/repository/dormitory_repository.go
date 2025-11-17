package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/your-org/go-backend-starter/internal/domain/entity"
	"github.com/your-org/go-backend-starter/internal/domain/repository"
	"github.com/your-org/go-backend-starter/internal/infrastructure/database"
	"gorm.io/gorm"
)

type dormitoryRepository struct {
	db *gorm.DB
}

// NewDormitoryRepository creates a new dormitory repository
func NewDormitoryRepository() repository.DormitoryRepository {
	return &dormitoryRepository{
		db: database.DB,
	}
}

func (r *dormitoryRepository) Create(ctx context.Context, dormitory *entity.Dormitory) error {
	return r.db.WithContext(ctx).Create(dormitory).Error
}

func (r *dormitoryRepository) GetByID(ctx context.Context, id uuid.UUID) (*entity.Dormitory, error) {
	var dormitory entity.Dormitory
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&dormitory).Error
	if err != nil {
		return nil, err
	}
	return &dormitory, nil
}

func (r *dormitoryRepository) Update(ctx context.Context, dormitory *entity.Dormitory) error {
	return r.db.WithContext(ctx).Save(dormitory).Error
}

func (r *dormitoryRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&entity.Dormitory{}, id).Error
}

func (r *dormitoryRepository) List(ctx context.Context, limit, offset int) ([]*entity.Dormitory, int64, error) {
	var dormitories []*entity.Dormitory
	var total int64

	err := r.db.WithContext(ctx).Model(&entity.Dormitory{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.WithContext(ctx).
		Limit(limit).
		Offset(offset).
		Find(&dormitories).Error

	return dormitories, total, err
}

func (r *dormitoryRepository) AssignToUser(ctx context.Context, userID, dormitoryID uuid.UUID) error {
	return r.db.WithContext(ctx).
		Create(&entity.UserDormitory{
			UserID:      userID,
			DormitoryID: dormitoryID,
		}).Error
}

func (r *dormitoryRepository) RemoveFromUser(ctx context.Context, userID, dormitoryID uuid.UUID) error {
	return r.db.WithContext(ctx).
		Where("user_id = ? AND dormitory_id = ?", userID, dormitoryID).
		Delete(&entity.UserDormitory{}).Error
}

func (r *dormitoryRepository) GetUserDormitories(ctx context.Context, userID uuid.UUID) ([]*entity.Dormitory, error) {
	var dormitories []*entity.Dormitory
	err := r.db.WithContext(ctx).
		Joins("JOIN user_dormitories ON user_dormitories.dormitory_id = dormitories.id").
		Where("user_dormitories.user_id = ?", userID).
		Find(&dormitories).Error
	return dormitories, err
}
