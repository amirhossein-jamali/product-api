package route

import (
	"product-api/controller"
	"product-api/model"

	"github.com/gorilla/mux"
)

// NewRouter initializes a new mux.Router and sets up the API routes for the product service.
// It takes a ProductRepository as input and wires it to the appropriate controller methods.
func NewRouter(repo model.ProductRepository) *mux.Router {
	// Create a new router instance
	r := mux.NewRouter()

	// Initialize the ProductController with the provided repository
	productController := controller.NewProductController(repo)

	// Define the route for fetching all products
	// Example: GET /products
	r.HandleFunc("/products", productController.GetProducts).Methods("GET")

	// Define the route for fetching a specific product by its ID
	// Example: GET /products/1
	// The route uses a regex to ensure the {id} is a numeric value.
	r.HandleFunc("/products/{id:[0-9]+}", productController.GetProductByID).Methods("GET")

	// Return the configured router
	return r
}
