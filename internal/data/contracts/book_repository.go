package contracts

import (
	uuid "github.com/satori/go.uuid"
	"go-template-microservice-v2/internal/data/entities"
)

type IBookRepository interface {
	AddBook(bookEntity entities.BookEntity) error
	GetBook(id uuid.UUID) (*entities.BookEntity, error)
	GetAllBook() ([]*entities.BookEntity, error)
	UpdateBook(bookEntity entities.BookEntity) error
}
