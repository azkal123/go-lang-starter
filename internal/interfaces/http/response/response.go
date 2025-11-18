package response

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// SuccessResponse represents a standardized success response
type SuccessResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// ErrorResponse represents a standardized error response
type ErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

// FieldError represents a single field validation error
type FieldError struct {
	Field   string `json:"field"`
	Rule    string `json:"rule"`
	Message string `json:"message"`
}

// ValidationErrorResponse represents a standardized validation error response
type ValidationErrorResponse struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Errors  []FieldError `json:"errors"`
}

// Success sends a standardized success response
func Success(c *gin.Context, statusCode int, data interface{}, message ...string) {
	msg := ""
	if len(message) > 0 {
		msg = message[0]
	}

	c.JSON(statusCode, SuccessResponse{
		Success: true,
		Message: msg,
		Data:    data,
	})
}

// Error sends a standardized error response
func Error(c *gin.Context, statusCode int, message string, errorDetail ...string) {
	errorDetailStr := ""
	if len(errorDetail) > 0 {
		errorDetailStr = errorDetail[0]
	}

	c.JSON(statusCode, ErrorResponse{
		Success: false,
		Message: message,
		Error:   errorDetailStr,
	})
}

// SuccessOK sends a 200 OK success response
func SuccessOK(c *gin.Context, data interface{}, message ...string) {
	Success(c, http.StatusOK, data, message...)
}

// SuccessCreated sends a 201 Created success response
func SuccessCreated(c *gin.Context, data interface{}, message ...string) {
	Success(c, http.StatusCreated, data, message...)
}

// SuccessNoContent sends a 204 No Content success response
func SuccessNoContent(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

// ErrorBadRequest sends a 400 Bad Request error response
func ErrorBadRequest(c *gin.Context, message string, errorDetail ...string) {
	if message == "" {
		message = "Bad request"
	}
	Error(c, http.StatusBadRequest, message, errorDetail...)
}

// ErrorValidation sends a 400 Bad Request response with structured validation errors
func ErrorValidation(c *gin.Context, err error) {
	var verrs validator.ValidationErrors
	if errors.As(err, &verrs) {
		fieldErrors := make([]FieldError, 0, len(verrs))
		for _, fe := range verrs {
			field := fe.Field()
			field = strings.TrimPrefix(field, "Create")
			field = strings.TrimPrefix(field, "Update")
			field = strings.TrimPrefix(field, "Register")
			field = strings.TrimPrefix(field, "Login")
			field = strings.TrimPrefix(field, "Request")
			field = strings.TrimPrefix(field, "Role")
			field = strings.TrimPrefix(field, "User")
			field = strings.TrimPrefix(field, "Dormitory")
			field = strings.TrimPrefix(field, "Permission")
			field = strings.TrimPrefix(field, "Location")
			field = strings.TrimPrefix(field, "DTO")
			field = strings.TrimPrefix(field, ".")
			field = strings.ToLower(field)

			message := fmt.Sprintf("%s is %s", field, fe.Tag())
			fieldErrors = append(fieldErrors, FieldError{
				Field:   field,
				Rule:    fe.Tag(),
				Message: message,
			})
		}

		c.JSON(http.StatusBadRequest, ValidationErrorResponse{
			Success: false,
			Message: "Validation failed",
			Errors:  fieldErrors,
		})
		return
	}

	// Fallback to generic bad request if it's not a validation error
	ErrorBadRequest(c, "Invalid request body", err.Error())
}

// ErrorUnauthorized sends a 401 Unauthorized error response
func ErrorUnauthorized(c *gin.Context, message string, errorDetail ...string) {
	if message == "" {
		message = "Unauthorized"
	}
	Error(c, http.StatusUnauthorized, message, errorDetail...)
}

// ErrorForbidden sends a 403 Forbidden error response
func ErrorForbidden(c *gin.Context, message string, errorDetail ...string) {
	if message == "" {
		message = "Forbidden"
	}
	Error(c, http.StatusForbidden, message, errorDetail...)
}

// ErrorNotFound sends a 404 Not Found error response
func ErrorNotFound(c *gin.Context, message string, errorDetail ...string) {
	if message == "" {
		message = "Not found"
	}
	Error(c, http.StatusNotFound, message, errorDetail...)
}

// ErrorConflict sends a 409 Conflict error response
func ErrorConflict(c *gin.Context, message string, errorDetail ...string) {
	if message == "" {
		message = "Conflict"
	}
	Error(c, http.StatusConflict, message, errorDetail...)
}

// ErrorInternalServer sends a 500 Internal Server Error response
func ErrorInternalServer(c *gin.Context, message string, errorDetail ...string) {
	if message == "" {
		message = "Internal server error"
	}
	Error(c, http.StatusInternalServerError, message, errorDetail...)
}
