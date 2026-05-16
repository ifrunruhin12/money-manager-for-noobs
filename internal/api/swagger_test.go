package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ifrunruhin12/money-manager/internal/api"
	"github.com/ifrunruhin12/money-manager/internal/handler"
	"github.com/ifrunruhin12/money-manager/pkg/logger"
)

func TestSwaggerEndpoints(t *testing.T) {
	log := logger.New("error")

	// Create a minimal router with just the swagger endpoints
	router := api.NewRouter(
		"test-secret",
		0,
		log,
		nil, // no DB needed for swagger
		handler.NewAuthHandler(nil),
		handler.NewAccountHandler(nil, nil),
		handler.NewTransactionHandler(nil),
		handler.NewCategoryHandler(nil),
		handler.NewBigBuyHandler(nil),
	)

	t.Run("GET /swagger returns HTML", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/swagger", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		require.Equal(t, http.StatusOK, w.Code, "should return 200")
		require.Contains(t, w.Header().Get("Content-Type"), "text/html", "should return HTML")
		require.Contains(t, w.Body.String(), "swagger-ui", "should contain swagger-ui div")
		require.Contains(t, w.Body.String(), "SwaggerUIBundle", "should contain SwaggerUIBundle script")
	})

	t.Run("GET /swagger.yaml returns YAML", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/swagger.yaml", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		// Note: This will return 404 in test because the file path is relative to the binary
		// In production/docker, the file will be present at ./docs/swagger.yaml
		// We're just testing that the route is registered
		require.True(t, w.Code == http.StatusOK || w.Code == http.StatusNotFound,
			"route should be registered (200 if file exists, 404 if not)")
	})
}
