package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ifrunruhin12/money-manager/internal/utils"
)

// BigBuyHandler handles big buy endpoints.
type BigBuyHandler struct{}

// NewBigBuyHandler creates a new BigBuyHandler.
func NewBigBuyHandler() *BigBuyHandler {
	return &BigBuyHandler{}
}

// Create handles POST /big-buys.
func (h *BigBuyHandler) Create(c *gin.Context) {
	userID, ok := utils.GetUserID(c)
	if !ok {
		return
	}
	_ = userID
	utils.WriteOK(c, http.StatusCreated, gin.H{})
}

// List handles GET /big-buys.
func (h *BigBuyHandler) List(c *gin.Context) {
	userID, ok := utils.GetUserID(c)
	if !ok {
		return
	}
	_ = userID
	utils.WriteOK(c, http.StatusOK, gin.H{"big_buys": []any{}})
}

// Update handles PATCH /big-buys/:id.
func (h *BigBuyHandler) Update(c *gin.Context) {
	userID, ok := utils.GetUserID(c)
	if !ok {
		return
	}
	_ = userID
	utils.WriteOK(c, http.StatusOK, gin.H{})
}

// Delete handles DELETE /big-buys/:id.
func (h *BigBuyHandler) Delete(c *gin.Context) {
	userID, ok := utils.GetUserID(c)
	if !ok {
		return
	}
	_ = userID
	utils.WriteOK(c, http.StatusOK, gin.H{})
}
