package repository

import (
	"test_task/models"

	"github.com/go-pg/pg/v10"
)

type AuthPQ struct {
	db *pg.DB
}

func NewAuthPQ(db *pg.DB) *AuthPQ {
	return &AuthPQ{db: db}
}

func (u AuthPQ) GetUser(phone string) (models.Users, error) {
	checkUs := new(models.Users)
	err := u.db.Model(checkUs).
		Where("users.number_phone = ?", phone).
		Select()
	if err == nil {
		return *checkUs, nil
	}
	if err.Error() == "pg: no rows in result set" {
		return *&models.Users{}, nil
	}
	return *&models.Users{}, err
}

func (u AuthPQ) CreateUser(phone string) error {
	newUs := &models.Users{
		Number_phone: phone,
	}
	_, err := u.db.Model(newUs).Insert()
	if err != nil {
		return err
	}
	return nil
}
