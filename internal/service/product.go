package service

import (
	"test_task/graph/model"
	"test_task/internal/repository"
)

type ProductService struct {
	repo repository.Product
}

func NewProductService(repo repository.Product) *ProductService {
	return &ProductService{repo: repo}
}

func (p ProductService) GetAll() ([]*model.Product, error) {
	return p.repo.GetAll()
}
