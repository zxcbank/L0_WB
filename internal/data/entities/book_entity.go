package entities

import (
	uuid "github.com/satori/go.uuid"
)

// BookEntity model
type BookEntity struct {
	Id      uuid.UUID `json:"id" gorm:"primaryKey"`
	Name    string    `json:"name"`
	Author  string    `json:"author"`
	Price   float64   `json:"price"`
	Enabled bool      `json:"enabled"`
}

// CreateBookEntity создать модель
func CreateBookEntity(name string, author string, price float64) BookEntity {
	return BookEntity{
		Name:    name,
		Author:  author,
		Price:   price,
		Id:      uuid.NewV4(),
		Enabled: true,
	}
}
