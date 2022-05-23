package services

import "example/models"

type ProductService interface {
	CreateProduct(*models.Product) error
	GetAll() ([]*models.Product, error)
	UpdateProduct(*models.Product) error
	DeleteProduct(*string) error
}
