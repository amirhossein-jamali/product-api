package controller

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"product-api/model"
	"product-api/util"
	"strconv"

	"github.com/gorilla/mux"
)

type ProductController struct {
	Repo model.ProductRepository
}

// NewProductController creates a new instance of ProductController.
func NewProductController(repo model.ProductRepository) *ProductController {
	return &ProductController{Repo: repo}
}

// GetProducts retrieves and responds with all products.
func (pc *ProductController) GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	products := pc.Repo.GetAllProducts()

	if err := json.NewEncoder(w).Encode(products); err != nil {
		log.Printf("Error encoding products: %v", err)
		util.WriteError(w, "Failed to retrieve products", http.StatusInternalServerError)
	}
}

// GetProductByID retrieves a single product by its ID.
func (pc *ProductController) GetProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		util.WriteError(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product, err := pc.Repo.GetProductByID(id)
	if err != nil {
		if errors.Is(err, model.ErrorNotFound) {
			util.WriteError(w, "Product not found", http.StatusNotFound)
		} else {
			log.Printf("Error retrieving product: %v", err)
			util.WriteError(w, "Failed to retrieve product", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(product); err != nil {
		log.Printf("Error encoding product: %v", err)
		util.WriteError(w, "Failed to encode product", http.StatusInternalServerError)
	}
}
