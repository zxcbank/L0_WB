package commands

import uuid "github.com/satori/go.uuid"

// BuyBookCommand - модель добавления книги в каталог
type BuyBookCommand struct {
	BookId uuid.UUID `json:"BookId"   validate:"required"`
}
