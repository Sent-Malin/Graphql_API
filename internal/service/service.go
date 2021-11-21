package service

import (
	"test_task/graph/model"
	"test_task/internal/repository"
)

type Authorization interface {
	SignIn(phone string) error
	GenerateToken(phone string) (string, error)
	ParseToken(accessToken string) (string, error)
}

type Product interface {
	GetAll() ([]*model.Product, error)
}

type Service struct {
	Authorization
	Product
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Product:       NewProductService(repos.Product),
	}
}
