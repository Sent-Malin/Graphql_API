package graph

import (
	"test_task/graph/model"
	"test_task/internal/service"
)

type Resolver struct {
	Authorization
	Product
	phonesCodes map[string]int
}

type Authorization interface {
	SignIn(phone string) error
	GenerateToken(phone string) (string, error)
	ParseToken(accessToken string) (string, error)
}

type Product interface {
	GetAll() ([]*model.Product, error)
}

func NewResolver(serv *service.Service) *Resolver {
	return &Resolver{
		Authorization: serv.Authorization,
		Product:       serv.Product,
		phonesCodes:   make(map[string]int),
	}
}
