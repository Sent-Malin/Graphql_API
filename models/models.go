package models

type Users struct {
	ID           int    `json:"id"`
	Number_phone string `json:"phone" pg:"number_phone"`
}
