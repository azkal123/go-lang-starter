package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/your-org/go-backend-starter/internal/application/dto"
	"github.com/your-org/go-backend-starter/internal/application/usecase"
	domainErrors "github.com/your-org/go-backend-starter/internal/domain/errors"
)

// UserHandler handles user management requests
type UserHandler struct {
	userUseCase *usecase.UserUseCase
}

// NewUserHandler creates a new user handler
func NewUserHandler(userUseCase *usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

// CreateUser handles user creation
// @Summary Create a new user
// @Description Create a new user (admin only)
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreateUserRequest true "Create user request"
// @Success 201 {object} dto.UserResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Router /api/users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.userUseCase.CreateUser(c.Request.Context(), req)
	if err != nil {
		switch err {
		case domainErrors.ErrUserAlreadyExists:
			c.JSON(http.StatusConflict, gin.H{"error": "user already exists"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		}
		return
	}

	c.JSON(http.StatusCreated, response)
}

// GetUser handles getting a user by ID
// @Summary Get user by ID
// @Description Get user details by ID
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "User ID"
// @Success 200 {object} dto.UserResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/users/{id} [get]
func (h *UserHandler) GetUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	response, err := h.userUseCase.GetUserByID(c.Request.Context(), id)
	if err != nil {
		switch err {
		case domainErrors.ErrUserNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get user"})
		}
		return
	}

	c.JSON(http.StatusOK, response)
}

// UpdateUser handles user update
// @Summary Update user
// @Description Update user details
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "User ID"
// @Param request body dto.UpdateUserRequest true "Update user request"
// @Success 200 {object} dto.UserResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/users/{id} [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	var req dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.userUseCase.UpdateUser(c.Request.Context(), id, req)
	if err != nil {
		switch err {
		case domainErrors.ErrUserNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		case domainErrors.ErrUserAlreadyExists:
			c.JSON(http.StatusConflict, gin.H{"error": "email already taken"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update user"})
		}
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteUser handles user deletion
// @Summary Delete user
// @Description Delete a user (soft delete)
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "User ID"
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/users/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	err = h.userUseCase.DeleteUser(c.Request.Context(), id)
	if err != nil {
		switch err {
		case domainErrors.ErrUserNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete user"})
		}
		return
	}

	c.Status(http.StatusNoContent)
}

// ListUsers handles listing users with pagination
// @Summary List users
// @Description Get paginated list of users
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Page size" default(10)
// @Success 200 {object} dto.ListUsersResponse
// @Failure 400 {object} map[string]string
// @Router /api/users [get]
func (h *UserHandler) ListUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	response, err := h.userUseCase.ListUsers(c.Request.Context(), page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list users"})
		return
	}

	c.JSON(http.StatusOK, response)
}
