package repository

import (
	"test_task/graph/model"
	"test_task/models"

	"github.com/go-pg/pg/v10"
)

type Authorization interface {
	GetUser(phone string) (models.Users, error)
	CreateUser(phone string) error
}

type Product interface {
	GetAll() ([]*model.Product, error)
}

type Repository struct {
	Authorization
	Product
}

func NewRepository(db *pg.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPQ(db),
		Product:       NewProduct(db),
	}
}
