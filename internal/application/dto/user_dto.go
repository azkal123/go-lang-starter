package dto

import "github.com/google/uuid"

// CreateUserRequest represents the request to create a user
type CreateUserRequest struct {
	Email    string   `json:"email" binding:"required,email"`
	Password string   `json:"password" binding:"required,min=6"`
	Name     string   `json:"name" binding:"required"`
	RoleIDs  []string `json:"role_ids,omitempty"`
}

// UpdateUserRequest represents the request to update a user
type UpdateUserRequest struct {
	Name     string   `json:"name,omitempty"`
	Email    string   `json:"email,omitempty"`
	IsActive *bool    `json:"is_active,omitempty"`
	RoleIDs  []string `json:"role_ids,omitempty"`
}

// UserResponse represents user data in responses
type UserResponse struct {
	ID        string   `json:"id"`
	Email     string   `json:"email"`
	Name      string   `json:"name"`
	IsActive  bool     `json:"is_active"`
	Roles     []string `json:"roles,omitempty"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

// ListUsersResponse represents paginated user list response
type ListUsersResponse struct {
	Users      []UserResponse `json:"users"`
	Total      int64          `json:"total"`
	Page       int            `json:"page"`
	PageSize   int            `json:"page_size"`
	TotalPages int            `json:"total_pages"`
}

// AssignRoleRequest represents the request to assign a role to a user
type AssignRoleRequest struct {
	UserID uuid.UUID `json:"user_id" binding:"required"`
	RoleID uuid.UUID `json:"role_id" binding:"required"`
}
