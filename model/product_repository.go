package model

import "errors"

var ErrorNotFound = errors.New("product not found")

type ProductRepository interface {
	GetAllProducts() []Product
	GetProductByID(id int) (*Product, error)
}
