package repository

import (
	"test_task/graph/model"

	"github.com/go-pg/pg/v10"
)

type ProductPQ struct {
	db *pg.DB
}

func NewProduct(db *pg.DB) *ProductPQ {
	return &ProductPQ{db: db}
}

func (p ProductPQ) GetAll() ([]*model.Product, error) {
	var products []*model.Product
	err := p.db.Model(&products).Select()

	if (err != nil) && (err.Error() != "pg: no rows in result set") {
		return nil, err
	}
	return products, nil
}
