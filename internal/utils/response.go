package utils

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ifrunruhin12/money-manager/internal/domain"
)

func BindJSON(c *gin.Context, obj any) bool {
	if err := c.ShouldBindJSON(obj); err != nil {
		AbortWithError(c, http.StatusBadRequest, "invalid request body")
		return false
	}
	return true
}

func WriteOK(c *gin.Context, status int, data any) {
	c.JSON(status, gin.H{"data": data, "ok": true})
}

func WriteError(c *gin.Context, status int, msg string) {
	c.JSON(status, gin.H{"error": msg, "ok": false})
}

func AbortWithError(c *gin.Context, status int, msg string) {
	c.AbortWithStatusJSON(status, gin.H{"error": msg, "ok": false})
}

func MapError(err error) (int, string) {
	switch {
	case errors.Is(err, domain.ErrNotFound):
		return http.StatusNotFound, "resource not found"
	case errors.Is(err, domain.ErrConflict):
		return http.StatusConflict, "resource already exists"
	case errors.Is(err, domain.ErrValidation):
		return http.StatusBadRequest, err.Error()
	case errors.Is(err, domain.ErrUnauthorized):
		return http.StatusUnauthorized, "invalid credentials"
	default:
		return http.StatusInternalServerError, "internal server error"
	}
}

func GetUserID(c *gin.Context) (string, bool) {
	val, exists := c.Get("user_id")
	if !exists {
		WriteError(c, http.StatusInternalServerError, "user_id not found in context")
		return "", false
	}
	userID, ok := val.(string)
	if !ok || userID == "" {
		WriteError(c, http.StatusInternalServerError, "invalid user_id in context")
		return "", false
	}
	return userID, true
}
