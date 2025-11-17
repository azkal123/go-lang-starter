package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/your-org/go-backend-starter/internal/domain/entity"
	"github.com/your-org/go-backend-starter/internal/domain/repository"
	"github.com/your-org/go-backend-starter/internal/infrastructure/database"
	"gorm.io/gorm"
)

type permissionRepository struct {
	db *gorm.DB
}

// NewPermissionRepository creates a new permission repository
func NewPermissionRepository() repository.PermissionRepository {
	return &permissionRepository{
		db: database.DB,
	}
}

func (r *permissionRepository) Create(ctx context.Context, permission *entity.Permission) error {
	return r.db.WithContext(ctx).Create(permission).Error
}

func (r *permissionRepository) GetByID(ctx context.Context, id uuid.UUID) (*entity.Permission, error) {
	var permission entity.Permission
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&permission).Error
	if err != nil {
		return nil, err
	}
	return &permission, nil
}

func (r *permissionRepository) GetBySlug(ctx context.Context, slug string) (*entity.Permission, error) {
	var permission entity.Permission
	err := r.db.WithContext(ctx).Where("slug = ?", slug).First(&permission).Error
	if err != nil {
		return nil, err
	}
	return &permission, nil
}

func (r *permissionRepository) Update(ctx context.Context, permission *entity.Permission) error {
	return r.db.WithContext(ctx).Save(permission).Error
}

func (r *permissionRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&entity.Permission{}, id).Error
}

func (r *permissionRepository) List(ctx context.Context, limit, offset int) ([]*entity.Permission, int64, error) {
	var permissions []*entity.Permission
	var total int64

	err := r.db.WithContext(ctx).Model(&entity.Permission{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.WithContext(ctx).
		Limit(limit).
		Offset(offset).
		Find(&permissions).Error

	return permissions, total, err
}
