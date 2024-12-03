package main

import (
	"fmt"
	"log"
	"net/http"
	"product-api/mocks"
	"product-api/model"
	"product-api/route"

	"github.com/stretchr/testify/mock"
)

func main() {
	// Mock repository for server testing
	mockRepo := &mocks.ProductRepository{}
	mockRepo.On("GetAllProducts").Return([]model.Product{
		{ID: 1, Name: "Product A"},
		{ID: 2, Name: "Product B"},
	})
	mockRepo.On("GetProductByID", mock.MatchedBy(func(id int) bool { return id == 1 })).Return(&model.Product{ID: 1, Name: "Product A"}, nil).Times(5)
	mockRepo.On("GetProductByID", mock.MatchedBy(func(id int) bool { return id == 2 })).Return(&model.Product{ID: 2, Name: "Product B"}, nil).Times(1)
	mockRepo.On("GetProductByID", mock.MatchedBy(func(id int) bool { return id != 1 && id != 2 })).Return(nil, model.ErrorNotFound)

	// Initialize router with mocked repository
	r := route.NewRouter(mockRepo)

	// Start server
	addr := ":2020"
	fmt.Printf("Server is running on http://localhost%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
