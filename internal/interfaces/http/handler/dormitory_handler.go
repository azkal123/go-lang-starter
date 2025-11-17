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

// DormitoryHandler handles dormitory management requests
type DormitoryHandler struct {
	dormitoryUseCase *usecase.DormitoryUseCase
}

// NewDormitoryHandler creates a new dormitory handler
func NewDormitoryHandler(dormitoryUseCase *usecase.DormitoryUseCase) *DormitoryHandler {
	return &DormitoryHandler{
		dormitoryUseCase: dormitoryUseCase,
	}
}

// CreateDormitory handles dormitory creation
// @Summary Create a new dormitory
// @Description Create a new dormitory
// @Tags dormitories
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreateDormitoryRequest true "Create dormitory request"
// @Success 201 {object} dto.DormitoryResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /api/dormitories [post]
func (h *DormitoryHandler) CreateDormitory(c *gin.Context) {
	var req dto.CreateDormitoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.dormitoryUseCase.CreateDormitory(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create dormitory"})
		return
	}

	c.JSON(http.StatusCreated, response)
}

// GetDormitory handles getting a dormitory by ID
// @Summary Get dormitory by ID
// @Description Get dormitory details by ID
// @Tags dormitories
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Dormitory ID"
// @Success 200 {object} dto.DormitoryResponse
// @Failure 400 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/dormitories/{id} [get]
func (h *DormitoryHandler) GetDormitory(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid dormitory id"})
		return
	}

	response, err := h.dormitoryUseCase.GetDormitoryByID(c.Request.Context(), id)
	if err != nil {
		switch err {
		case domainErrors.ErrDormitoryNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": "dormitory not found"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get dormitory"})
		}
		return
	}

	c.JSON(http.StatusOK, response)
}

// UpdateDormitory handles dormitory update
// @Summary Update dormitory
// @Description Update dormitory details
// @Tags dormitories
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Dormitory ID"
// @Param request body dto.UpdateDormitoryRequest true "Update dormitory request"
// @Success 200 {object} dto.DormitoryResponse
// @Failure 400 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/dormitories/{id} [put]
func (h *DormitoryHandler) UpdateDormitory(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid dormitory id"})
		return
	}

	var req dto.UpdateDormitoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.dormitoryUseCase.UpdateDormitory(c.Request.Context(), id, req)
	if err != nil {
		switch err {
		case domainErrors.ErrDormitoryNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": "dormitory not found"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update dormitory"})
		}
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteDormitory handles dormitory deletion
// @Summary Delete dormitory
// @Description Delete a dormitory (soft delete)
// @Tags dormitories
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Dormitory ID"
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/dormitories/{id} [delete]
func (h *DormitoryHandler) DeleteDormitory(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid dormitory id"})
		return
	}

	err = h.dormitoryUseCase.DeleteDormitory(c.Request.Context(), id)
	if err != nil {
		switch err {
		case domainErrors.ErrDormitoryNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": "dormitory not found"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete dormitory"})
		}
		return
	}

	c.Status(http.StatusNoContent)
}

// ListDormitories handles listing dormitories with pagination
// @Summary List dormitories
// @Description Get paginated list of dormitories
// @Tags dormitories
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Page size" default(10)
// @Success 200 {object} dto.ListDormitoriesResponse
// @Failure 400 {object} map[string]string
// @Router /api/dormitories [get]
func (h *DormitoryHandler) ListDormitories(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	response, err := h.dormitoryUseCase.ListDormitories(c.Request.Context(), page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list dormitories"})
		return
	}

	c.JSON(http.StatusOK, response)
}
