package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"product-api/mocks"
	"product-api/model"
	"product-api/util"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func setupControllerWithMock() (*ProductController, *mocks.ProductRepository) {
	mockRepo := &mocks.ProductRepository{}
	controller := NewProductController(mockRepo)
	return controller, mockRepo
}

func TestGetProducts(t *testing.T) {
	controller, mockRepo := setupControllerWithMock()

	mockRepo.On("GetAllProducts").Return([]model.Product{
		{ID: 1, Name: "Product A"},
		{ID: 2, Name: "Product B"},
	}).Once()

	req, _ := http.NewRequest("GET", "/products", nil)
	rr := httptest.NewRecorder()

	controller.GetProducts(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var products []model.Product
	err := json.Unmarshal(rr.Body.Bytes(), &products)
	assert.NoError(t, err)
	assert.Len(t, products, 2)

	mockRepo.AssertExpectations(t)
}

func TestGetProductByID_Success(t *testing.T) {
	controller, mockRepo := setupControllerWithMock()
	expectedProduct := &model.Product{ID: 1, Name: "Product A"}

	mockRepo.On("GetProductByID", 1).Return(expectedProduct, nil).Once()

	req, _ := http.NewRequest("GET", "/products/1", nil)
	rr := httptest.NewRecorder()

	// Simulate mux variable setup
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	controller.GetProductByID(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var product model.Product
	err := json.Unmarshal(rr.Body.Bytes(), &product)
	assert.NoError(t, err)
	assert.Equal(t, expectedProduct, &product)

	mockRepo.AssertExpectations(t)
}

func TestGetProductByID_NotFound(t *testing.T) {
	controller, mockRepo := setupControllerWithMock()

	mockRepo.On("GetProductByID", 999).Return(nil, model.ErrorNotFound).Once()

	req, _ := http.NewRequest("GET", "/products/999", nil)
	rr := httptest.NewRecorder()

	// Simulate mux variable setup
	req = mux.SetURLVars(req, map[string]string{"id": "999"})
	controller.GetProductByID(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)

	var apiError util.APIError
	err := json.Unmarshal(rr.Body.Bytes(), &apiError)
	assert.NoError(t, err)
	assert.Equal(t, "Product not found", apiError.Message)

	mockRepo.AssertExpectations(t)
}

func TestGetProductByID_InvalidID(t *testing.T) {
	controller, _ := setupControllerWithMock()

	req, _ := http.NewRequest("GET", "/products/abc", nil)
	rr := httptest.NewRecorder()

	// Simulate mux variable setup
	req = mux.SetURLVars(req, map[string]string{"id": "abc"})
	controller.GetProductByID(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)

	var apiError util.APIError
	err := json.Unmarshal(rr.Body.Bytes(), &apiError)
	assert.NoError(t, err)
	assert.Equal(t, "Invalid product ID", apiError.Message)
}
