package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/ifrunruhin12/money-manager/internal/domain"
	"github.com/ifrunruhin12/money-manager/internal/service"
)

// MockTransactionService is a mock implementation of service.TransactionService
type MockTransactionService struct {
	mock.Mock
}

func (m *MockTransactionService) Create(ctx context.Context, tx domain.Transaction) (*domain.Transaction, error) {
	args := m.Called(ctx, tx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Transaction), args.Error(1)
}

func (m *MockTransactionService) Override(ctx context.Context, originalID, userID string, amount int, note string) (*domain.Transaction, error) {
	args := m.Called(ctx, originalID, userID, amount, note)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Transaction), args.Error(1)
}

func (m *MockTransactionService) Skip(ctx context.Context, txID, userID string) error {
	args := m.Called(ctx, txID, userID)
	return args.Error(0)
}

func (m *MockTransactionService) Restore(ctx context.Context, txID, userID string) error {
	args := m.Called(ctx, txID, userID)
	return args.Error(0)
}

func (m *MockTransactionService) ListByDateRange(ctx context.Context, userID string, from, to time.Time, limit int, cursorDate time.Time, cursorID string) ([]domain.Transaction, *service.Cursor, error) {
	args := m.Called(ctx, userID, from, to, limit, cursorDate, cursorID)
	if args.Get(0) == nil {
		return nil, nil, args.Error(2)
	}
	cursor, _ := args.Get(1).(*service.Cursor)
	return args.Get(0).([]domain.Transaction), cursor, args.Error(2)
}

func (m *MockTransactionService) GetHistory(ctx context.Context, txID, userID string) ([]domain.Transaction, error) {
	args := m.Called(ctx, txID, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]domain.Transaction), args.Error(1)
}

// MockCategoryService is a mock implementation of service.CategoryService
type MockCategoryService struct {
	mock.Mock
}

func (m *MockCategoryService) Create(ctx context.Context, userID, name string) (*domain.Category, error) {
	args := m.Called(ctx, userID, name)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Category), args.Error(1)
}

func (m *MockCategoryService) Update(ctx context.Context, id, userID, name string) (*domain.Category, error) {
	args := m.Called(ctx, id, userID, name)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Category), args.Error(1)
}

func (m *MockCategoryService) Delete(ctx context.Context, id, userID string) error {
	args := m.Called(ctx, id, userID)
	return args.Error(0)
}

func (m *MockCategoryService) List(ctx context.Context, userID string) ([]domain.Category, error) {
	args := m.Called(ctx, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]domain.Category), args.Error(1)
}

// MockBigBuyService is a mock implementation of service.BigBuyService
type MockBigBuyService struct {
	mock.Mock
}

func (m *MockBigBuyService) Create(ctx context.Context, bigBuy domain.BigBuy) (*domain.BigBuy, error) {
	args := m.Called(ctx, bigBuy)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.BigBuy), args.Error(1)
}

func (m *MockBigBuyService) Update(ctx context.Context, bigBuy domain.BigBuy) (*domain.BigBuy, error) {
	args := m.Called(ctx, bigBuy)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.BigBuy), args.Error(1)
}

func (m *MockBigBuyService) Delete(ctx context.Context, id, userID string) error {
	args := m.Called(ctx, id, userID)
	return args.Error(0)
}

func (m *MockBigBuyService) ListByMonth(ctx context.Context, userID string, year, month int) ([]domain.BigBuy, error) {
	args := m.Called(ctx, userID, year, month)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]domain.BigBuy), args.Error(1)
}

// setupTestRouter creates a test Gin router with test mode enabled
func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	return gin.New()
}

// injectUserID is a middleware that injects a test user ID into the context
func injectUserID(userID string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("user_id", userID)
		c.Next()
	}
}

// Property 39: HTTP 400 for Malformed JSON
// Validates: Requirements 13.2
// Tests that all handlers return HTTP 400 with descriptive error messages when receiving malformed JSON

func TestProperty39_MalformedJSON_TransactionCreate(t *testing.T) {
	router := setupTestRouter()
	mockService := new(MockTransactionService)
	handler := NewTransactionHandler(mockService)

	router.POST("/transactions", injectUserID("test-user"), handler.Create)

	// Test with invalid JSON (missing closing brace)
	malformedJSON := `{"category_id": "cat-1", "amount": 100`

	req := httptest.NewRequest(http.MethodPost, "/transactions", bytes.NewBufferString(malformedJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.False(t, response["ok"].(bool))
	assert.Contains(t, response["error"].(string), "invalid request body")
}

func TestProperty39_MalformedJSON_CategoryCreate(t *testing.T) {
	router := setupTestRouter()
	mockService := new(MockCategoryService)
	handler := NewCategoryHandler(mockService)

	router.POST("/categories", injectUserID("test-user"), handler.Create)

	// Test with invalid JSON (not a JSON object)
	malformedJSON := `not json at all`

	req := httptest.NewRequest(http.MethodPost, "/categories", bytes.NewBufferString(malformedJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.False(t, response["ok"].(bool))
	assert.NotEmpty(t, response["error"])
}

func TestProperty39_MalformedJSON_BigBuyCreate(t *testing.T) {
	router := setupTestRouter()
	mockService := new(MockBigBuyService)
	handler := NewBigBuyHandler(mockService)

	router.POST("/big-buys", injectUserID("test-user"), handler.Create)

	// Test with invalid JSON (array instead of object)
	malformedJSON := `["title", "amount"]`

	req := httptest.NewRequest(http.MethodPost, "/big-buys", bytes.NewBufferString(malformedJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.False(t, response["ok"].(bool))
	assert.Contains(t, response["error"].(string), "invalid request body")
}

func TestProperty39_MalformedJSON_TransactionOverride(t *testing.T) {
	router := setupTestRouter()
	mockService := new(MockTransactionService)
	handler := NewTransactionHandler(mockService)

	router.PATCH("/transactions/:id/override", injectUserID("test-user"), handler.Override)

	// Test with invalid JSON (wrong type for amount field)
	malformedJSON := `{"amount": "not-a-number", "note": "test"}`

	req := httptest.NewRequest(http.MethodPatch, "/transactions/tx-1/override", bytes.NewBufferString(malformedJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.False(t, response["ok"].(bool))
	assert.Contains(t, response["error"].(string), "invalid request body")
}

func TestProperty39_MalformedJSON_BigBuyUpdate(t *testing.T) {
	router := setupTestRouter()
	mockService := new(MockBigBuyService)
	handler := NewBigBuyHandler(mockService)

	router.PATCH("/big-buys/:id", injectUserID("test-user"), handler.Update)

	// Test with empty JSON body
	malformedJSON := ``

	req := httptest.NewRequest(http.MethodPatch, "/big-buys/bb-1", bytes.NewBufferString(malformedJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.False(t, response["ok"].(bool))
	assert.NotEmpty(t, response["error"])
}

// Property 40: HTTP 404 for Missing Resources
// Validates: Requirements 13.3
// Tests that handlers return HTTP 404 with descriptive error messages when resources don't exist

func TestProperty40_NotFound_TransactionSkip(t *testing.T) {
	router := setupTestRouter()
	mockService := new(MockTransactionService)
	handler := NewTransactionHandler(mockService)

	router.PATCH("/transactions/:id/skip", injectUserID("test-user"), handler.Skip)

	// Mock service to return ErrNotFound
	mockService.On("Skip", mock.Anything, "nonexistent-tx", "test-user").Return(domain.ErrNotFound)

	req := httptest.NewRequest(http.MethodPatch, "/transactions/nonexistent-tx/skip", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.False(t, response["ok"].(bool))
	assert.NotEmpty(t, response["error"])
	
	mockService.AssertExpectations(t)
}

func TestProperty40_NotFound_TransactionRestore(t *testing.T) {
	router := setupTestRouter()
	mockService := new(MockTransactionService)
	handler := NewTransactionHandler(mockService)

	router.PATCH("/transactions/:id/restore", injectUserID("test-user"), handler.Restore)

	// Mock service to return ErrNotFound
	mockService.On("Restore", mock.Anything, "nonexistent-tx", "test-user").Return(domain.ErrNotFound)

	req := httptest.NewRequest(http.MethodPatch, "/transactions/nonexistent-tx/restore", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.False(t, response["ok"].(bool))
	assert.NotEmpty(t, response["error"])
	
	mockService.AssertExpectations(t)
}

func TestProperty40_NotFound_TransactionOverride(t *testing.T) {
	router := setupTestRouter()
	mockService := new(MockTransactionService)
	handler := NewTransactionHandler(mockService)

	router.PATCH("/transactions/:id/override", injectUserID("test-user"), handler.Override)

	// Mock service to return ErrNotFound
	mockService.On("Override", mock.Anything, "nonexistent-tx", "test-user", 500, "test note").Return(nil, domain.ErrNotFound)

	reqBody := `{"amount": 500, "note": "test note"}`
	req := httptest.NewRequest(http.MethodPatch, "/transactions/nonexistent-tx/override", bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.False(t, response["ok"].(bool))
	assert.NotEmpty(t, response["error"])
	
	mockService.AssertExpectations(t)
}

func TestProperty40_NotFound_TransactionHistory(t *testing.T) {
	router := setupTestRouter()
	mockService := new(MockTransactionService)
	handler := NewTransactionHandler(mockService)

	router.GET("/transactions/:id/history", injectUserID("test-user"), handler.GetHistory)

	// Mock service to return ErrNotFound
	mockService.On("GetHistory", mock.Anything, "nonexistent-tx", "test-user").Return(nil, domain.ErrNotFound)

	req := httptest.NewRequest(http.MethodGet, "/transactions/nonexistent-tx/history", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.False(t, response["ok"].(bool))
	assert.NotEmpty(t, response["error"])
	
	mockService.AssertExpectations(t)
}

func TestProperty40_NotFound_CategoryUpdate(t *testing.T) {
	router := setupTestRouter()
	mockService := new(MockCategoryService)
	handler := NewCategoryHandler(mockService)

	router.PATCH("/categories/:id", injectUserID("test-user"), handler.Update)

	// Mock service to return ErrNotFound
	mockService.On("Update", mock.Anything, "nonexistent-cat", "test-user", "New Name").Return(nil, domain.ErrNotFound)

	reqBody := `{"name": "New Name"}`
	req := httptest.NewRequest(http.MethodPatch, "/categories/nonexistent-cat", bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.False(t, response["ok"].(bool))
	assert.NotEmpty(t, response["error"])
	
	mockService.AssertExpectations(t)
}

func TestProperty40_NotFound_CategoryDelete(t *testing.T) {
	router := setupTestRouter()
	mockService := new(MockCategoryService)
	handler := NewCategoryHandler(mockService)

	router.DELETE("/categories/:id", injectUserID("test-user"), handler.Delete)

	// Mock service to return ErrNotFound
	mockService.On("Delete", mock.Anything, "nonexistent-cat", "test-user").Return(domain.ErrNotFound)

	req := httptest.NewRequest(http.MethodDelete, "/categories/nonexistent-cat", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.False(t, response["ok"].(bool))
	assert.NotEmpty(t, response["error"])
	
	mockService.AssertExpectations(t)
}

func TestProperty40_NotFound_BigBuyUpdate(t *testing.T) {
	router := setupTestRouter()
	mockService := new(MockBigBuyService)
	handler := NewBigBuyHandler(mockService)

	router.PATCH("/big-buys/:id", injectUserID("test-user"), handler.Update)

	// Mock service to return ErrNotFound
	mockService.On("Update", mock.Anything, mock.AnythingOfType("domain.BigBuy")).Return(nil, domain.ErrNotFound)

	reqBody := `{"title": "Test", "amount": -1000, "category_id": "cat-1", "date": "2024-01-01T00:00:00Z", "note": ""}`
	req := httptest.NewRequest(http.MethodPatch, "/big-buys/nonexistent-bb", bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.False(t, response["ok"].(bool))
	assert.NotEmpty(t, response["error"])
	
	mockService.AssertExpectations(t)
}

func TestProperty40_NotFound_BigBuyDelete(t *testing.T) {
	router := setupTestRouter()
	mockService := new(MockBigBuyService)
	handler := NewBigBuyHandler(mockService)

	router.DELETE("/big-buys/:id", injectUserID("test-user"), handler.Delete)

	// Mock service to return ErrNotFound
	mockService.On("Delete", mock.Anything, "nonexistent-bb", "test-user").Return(domain.ErrNotFound)

	req := httptest.NewRequest(http.MethodDelete, "/big-buys/nonexistent-bb", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.False(t, response["ok"].(bool))
	assert.NotEmpty(t, response["error"])
	
	mockService.AssertExpectations(t)
}

// Additional edge case tests for comprehensive coverage

func TestMalformedJSON_MissingRequiredFields(t *testing.T) {
	router := setupTestRouter()
	mockService := new(MockTransactionService)
	handler := NewTransactionHandler(mockService)

	router.POST("/transactions", injectUserID("test-user"), handler.Create)

	// Missing required fields (amount and date)
	incompleteJSON := `{"category_id": "cat-1"}`

	req := httptest.NewRequest(http.MethodPost, "/transactions", bytes.NewBufferString(incompleteJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.False(t, response["ok"].(bool))
	assert.Contains(t, response["error"].(string), "invalid request body")
}

func TestMalformedJSON_InvalidDateFormat(t *testing.T) {
	router := setupTestRouter()
	mockService := new(MockTransactionService)
	handler := NewTransactionHandler(mockService)

	router.POST("/transactions", injectUserID("test-user"), handler.Create)

	// Invalid date format
	invalidDateJSON := `{"category_id": "cat-1", "amount": 100, "date": "not-a-date"}`

	req := httptest.NewRequest(http.MethodPost, "/transactions", bytes.NewBufferString(invalidDateJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.False(t, response["ok"].(bool))
	assert.Contains(t, response["error"].(string), "invalid request body")
}

func TestQueryParameterValidation_MissingRequired(t *testing.T) {
	router := setupTestRouter()
	mockService := new(MockTransactionService)
	handler := NewTransactionHandler(mockService)

	router.GET("/transactions", injectUserID("test-user"), handler.List)

	// Missing required 'from' query parameter
	req := httptest.NewRequest(http.MethodGet, "/transactions?to=2024-01-31T23:59:59Z", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.False(t, response["ok"].(bool))
	assert.Contains(t, response["error"].(string), "from query parameter is required")
}

func TestQueryParameterValidation_InvalidFormat(t *testing.T) {
	router := setupTestRouter()
	mockService := new(MockTransactionService)
	handler := NewTransactionHandler(mockService)

	router.GET("/transactions", injectUserID("test-user"), handler.List)

	// Invalid date format in query parameter
	req := httptest.NewRequest(http.MethodGet, "/transactions?from=invalid-date&to=2024-01-31T23:59:59Z", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.False(t, response["ok"].(bool))
	assert.Contains(t, response["error"].(string), "invalid from date format")
}

func TestBigBuyList_InvalidMonthFormat(t *testing.T) {
	router := setupTestRouter()
	mockService := new(MockBigBuyService)
	handler := NewBigBuyHandler(mockService)

	router.GET("/big-buys", injectUserID("test-user"), handler.List)

	// Invalid month format
	req := httptest.NewRequest(http.MethodGet, "/big-buys?month=invalid", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.False(t, response["ok"].(bool))
	assert.Contains(t, response["error"].(string), "invalid month format")
}

func TestBigBuyList_MissingMonthParameter(t *testing.T) {
	router := setupTestRouter()
	mockService := new(MockBigBuyService)
	handler := NewBigBuyHandler(mockService)

	router.GET("/big-buys", injectUserID("test-user"), handler.List)

	// Missing month parameter
	req := httptest.NewRequest(http.MethodGet, "/big-buys", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.False(t, response["ok"].(bool))
	assert.Contains(t, response["error"].(string), "month query parameter is required")
}
