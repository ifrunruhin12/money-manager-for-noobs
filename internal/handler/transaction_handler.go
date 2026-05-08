package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ifrunruhin12/money-manager/internal/utils"
)

// TransactionHandler handles transaction endpoints.
type TransactionHandler struct{}

// NewTransactionHandler creates a new TransactionHandler.
func NewTransactionHandler() *TransactionHandler {
	return &TransactionHandler{}
}

// Create handles POST /transactions.
func (h *TransactionHandler) Create(c *gin.Context) {
	userID, ok := utils.GetUserID(c)
	if !ok {
		return
	}
	_ = userID
	utils.WriteOK(c, http.StatusCreated, gin.H{})
}

// List handles GET /transactions.
func (h *TransactionHandler) List(c *gin.Context) {
	userID, ok := utils.GetUserID(c)
	if !ok {
		return
	}
	_ = userID
	utils.WriteOK(c, http.StatusOK, gin.H{"transactions": []any{}})
}

// Override handles PATCH /transactions/:id/override.
func (h *TransactionHandler) Override(c *gin.Context) {
	userID, ok := utils.GetUserID(c)
	if !ok {
		return
	}
	_ = userID
	utils.WriteOK(c, http.StatusOK, gin.H{})
}

// Skip handles PATCH /transactions/:id/skip.
func (h *TransactionHandler) Skip(c *gin.Context) {
	userID, ok := utils.GetUserID(c)
	if !ok {
		return
	}
	_ = userID
	utils.WriteOK(c, http.StatusOK, gin.H{})
}

// Restore handles PATCH /transactions/:id/restore.
func (h *TransactionHandler) Restore(c *gin.Context) {
	userID, ok := utils.GetUserID(c)
	if !ok {
		return
	}
	_ = userID
	utils.WriteOK(c, http.StatusOK, gin.H{})
}

// GetHistory handles GET /transactions/:id/history.
func (h *TransactionHandler) GetHistory(c *gin.Context) {
	userID, ok := utils.GetUserID(c)
	if !ok {
		return
	}
	_ = userID
	utils.WriteOK(c, http.StatusOK, gin.H{"history": []any{}})
}
