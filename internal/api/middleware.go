package api

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ifrunruhin12/money-manager/internal/utils"
)

const ContextUserIDKey = "user_id"

func AuthMiddleware(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.AbortWithError(c, http.StatusUnauthorized, "unauthorized")
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			utils.AbortWithError(c, http.StatusUnauthorized, "unauthorized")
			return
		}

		tokenString := strings.TrimSpace(parts[1])
		if tokenString == "" {
			utils.AbortWithError(c, http.StatusUnauthorized, "unauthorized")
			return
		}

		userID, err := utils.ParseToken(tokenString, jwtSecret)
		if err != nil {
			utils.AbortWithError(c, http.StatusUnauthorized, "unauthorized")
			return
		}

		c.Set(ContextUserIDKey, userID)
		c.Next()
	}
}

func TimeoutMiddleware(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
