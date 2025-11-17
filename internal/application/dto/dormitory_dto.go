package dto

// CreateDormitoryRequest represents the request to create a dormitory
type CreateDormitoryRequest struct {
	Name        string `json:"name" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Description string `json:"description"`
	Capacity    int    `json:"capacity" binding:"required,min=1"`
}

// UpdateDormitoryRequest represents the request to update a dormitory
type UpdateDormitoryRequest struct {
	Name        string `json:"name,omitempty"`
	Address     string `json:"address,omitempty"`
	Description string `json:"description,omitempty"`
	Capacity    *int   `json:"capacity,omitempty"`
	IsActive    *bool  `json:"is_active,omitempty"`
}

// DormitoryResponse represents dormitory data in responses
type DormitoryResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Description string `json:"description"`
	Capacity    int    `json:"capacity"`
	IsActive    bool   `json:"is_active"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// ListDormitoriesResponse represents paginated dormitory list response
type ListDormitoriesResponse struct {
	Dormitories []DormitoryResponse `json:"dormitories"`
	Total       int64               `json:"total"`
	Page        int                 `json:"page"`
	PageSize    int                 `json:"page_size"`
	TotalPages  int                 `json:"total_pages"`
}
