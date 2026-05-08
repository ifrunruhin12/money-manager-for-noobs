package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ifrunruhin12/money-manager/internal/utils"
)

// AccountHandler handles account and balance endpoints.
type AccountHandler struct{}

// NewAccountHandler creates a new AccountHandler.
func NewAccountHandler() *AccountHandler {
	return &AccountHandler{}
}

// GetBalance handles GET /balance.
func (h *AccountHandler) GetBalance(c *gin.Context) {
	userID, ok := utils.GetUserID(c)
	if !ok {
		return
	}
	_ = userID
	utils.WriteOK(c, http.StatusOK, gin.H{"balance": 0})
}

// UpdateStartingBalance handles PATCH /account/balance.
func (h *AccountHandler) UpdateStartingBalance(c *gin.Context) {
	userID, ok := utils.GetUserID(c)
	if !ok {
		return
	}
	_ = userID
	utils.WriteOK(c, http.StatusOK, gin.H{})
}

// UpdateTimezone handles PATCH /account/timezone.
func (h *AccountHandler) UpdateTimezone(c *gin.Context) {
	userID, ok := utils.GetUserID(c)
	if !ok {
		return
	}
	_ = userID
	utils.WriteOK(c, http.StatusOK, gin.H{})
}

// Reconcile handles POST /account/reconcile.
func (h *AccountHandler) Reconcile(c *gin.Context) {
	userID, ok := utils.GetUserID(c)
	if !ok {
		return
	}
	_ = userID
	utils.WriteOK(c, http.StatusOK, gin.H{"reconciled": true, "mismatch": false})
}
