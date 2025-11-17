package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/your-org/go-backend-starter/internal/domain/entity"
	"github.com/your-org/go-backend-starter/internal/domain/repository"
	"github.com/your-org/go-backend-starter/internal/infrastructure/database"
	"gorm.io/gorm"
)

type roleRepository struct {
	db *gorm.DB
}

// NewRoleRepository creates a new role repository
func NewRoleRepository() repository.RoleRepository {
	return &roleRepository{
		db: database.DB,
	}
}

func (r *roleRepository) Create(ctx context.Context, role *entity.Role) error {
	return r.db.WithContext(ctx).Create(role).Error
}

func (r *roleRepository) GetByID(ctx context.Context, id uuid.UUID) (*entity.Role, error) {
	var role entity.Role
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&role).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *roleRepository) GetBySlug(ctx context.Context, slug string) (*entity.Role, error) {
	var role entity.Role
	err := r.db.WithContext(ctx).Where("slug = ?", slug).First(&role).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *roleRepository) Update(ctx context.Context, role *entity.Role) error {
	return r.db.WithContext(ctx).Save(role).Error
}

func (r *roleRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&entity.Role{}, id).Error
}

func (r *roleRepository) List(ctx context.Context, limit, offset int) ([]*entity.Role, int64, error) {
	var roles []*entity.Role
	var total int64

	err := r.db.WithContext(ctx).Model(&entity.Role{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.WithContext(ctx).
		Limit(limit).
		Offset(offset).
		Find(&roles).Error

	return roles, total, err
}

func (r *roleRepository) GetWithPermissions(ctx context.Context, id uuid.UUID) (*entity.Role, error) {
	var role entity.Role
	err := r.db.WithContext(ctx).
		Preload("Permissions").
		Where("id = ?", id).
		First(&role).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *roleRepository) AssignPermission(ctx context.Context, roleID, permissionID uuid.UUID) error {
	return r.db.WithContext(ctx).
		Create(&entity.RolePermission{
			RoleID:       roleID,
			PermissionID: permissionID,
		}).Error
}

func (r *roleRepository) RemovePermission(ctx context.Context, roleID, permissionID uuid.UUID) error {
	return r.db.WithContext(ctx).
		Where("role_id = ? AND permission_id = ?", roleID, permissionID).
		Delete(&entity.RolePermission{}).Error
}
