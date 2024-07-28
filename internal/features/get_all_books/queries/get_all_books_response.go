package queries

import uuid "github.com/satori/go.uuid"

type GetAllBooksResponse struct {
	Books []GetAllBooksResponseItem `json:"books,omitempty"`
}

type GetAllBooksResponseItem struct {
	Id      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Author  string    `json:"author"`
	Price   float64   `json:"price"`
	Enabled bool      `json:"enabled"`
}
