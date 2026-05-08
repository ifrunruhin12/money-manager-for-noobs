package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ifrunruhin12/money-manager/internal/utils"
)

// CategoryHandler handles category endpoints.
type CategoryHandler struct{}

// NewCategoryHandler creates a new CategoryHandler.
func NewCategoryHandler() *CategoryHandler {
	return &CategoryHandler{}
}

// Create handles POST /categories.
func (h *CategoryHandler) Create(c *gin.Context) {
	userID, ok := utils.GetUserID(c)
	if !ok {
		return
	}
	_ = userID
	utils.WriteOK(c, http.StatusCreated, gin.H{})
}

// List handles GET /categories.
func (h *CategoryHandler) List(c *gin.Context) {
	userID, ok := utils.GetUserID(c)
	if !ok {
		return
	}
	_ = userID
	utils.WriteOK(c, http.StatusOK, gin.H{"categories": []any{}})
}

// Update handles PATCH /categories/:id.
func (h *CategoryHandler) Update(c *gin.Context) {
	userID, ok := utils.GetUserID(c)
	if !ok {
		return
	}
	_ = userID
	utils.WriteOK(c, http.StatusOK, gin.H{})
}

// Delete handles DELETE /categories/:id.
func (h *CategoryHandler) Delete(c *gin.Context) {
	userID, ok := utils.GetUserID(c)
	if !ok {
		return
	}
	_ = userID
	utils.WriteOK(c, http.StatusOK, gin.H{})
}
