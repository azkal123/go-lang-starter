package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/your-org/go-backend-starter/internal/application/dto"
	"github.com/your-org/go-backend-starter/internal/application/usecase"
	domainErrors "github.com/your-org/go-backend-starter/internal/domain/errors"
)

// AuthHandler handles authentication requests
type AuthHandler struct {
	authUseCase *usecase.AuthUseCase
}

// NewAuthHandler creates a new auth handler
func NewAuthHandler(authUseCase *usecase.AuthUseCase) *AuthHandler {
	return &AuthHandler{
		authUseCase: authUseCase,
	}
}

// Register handles user registration
// @Summary Register a new user
// @Description Register a new user account
// @Tags auth
// @Accept json
// @Produce json
// @Param request body dto.RegisterRequest true "Registration request"
// @Success 201 {object} dto.AuthResponse
// @Failure 400 {object} map[string]string
// @Failure 409 {object} map[string]string
// @Router /api/auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

		response, err := h.authUseCase.Register(c.Request.Context(), req)
	if err != nil {
		switch err {
		case domainErrors.ErrUserAlreadyExists:
			c.JSON(http.StatusConflict, gin.H{"error": "user already exists"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to register user"})
		}
		return
	}

	c.JSON(http.StatusCreated, response)
}

// Login handles user login
// @Summary Login user
// @Description Authenticate user and return tokens
// @Tags auth
// @Accept json
// @Produce json
// @Param request body dto.LoginRequest true "Login request"
// @Success 200 {object} dto.AuthResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /api/auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

		response, err := h.authUseCase.Login(c.Request.Context(), req)
	if err != nil {
		switch err {
		case domainErrors.ErrInvalidCredentials, domainErrors.ErrUserInactive:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to login"})
		}
		return
	}

	c.JSON(http.StatusOK, response)
}

// RefreshToken handles token refresh
// @Summary Refresh access token
// @Description Generate new access token using refresh token
// @Tags auth
// @Accept json
// @Produce json
// @Param request body dto.RefreshTokenRequest true "Refresh token request"
// @Success 200 {object} dto.AuthResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /api/auth/refresh [post]
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	var req dto.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

		response, err := h.authUseCase.RefreshToken(c.Request.Context(), req)
	if err != nil {
		switch err {
		case domainErrors.ErrInvalidToken, domainErrors.ErrTokenExpired:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to refresh token"})
		}
		return
	}

	c.JSON(http.StatusOK, response)
}
