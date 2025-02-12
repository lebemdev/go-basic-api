package services

import "go-basic-api/internal/domain"

type ProductService struct {
}

func (s ProductService) Create(product domain.Product) (uint, error) {
	return 0, nil
}

func (s ProductService) GetProducts() (string, error) {
	return "products", nil
}
