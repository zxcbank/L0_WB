package commands

import uuid "github.com/satori/go.uuid"

type AddBookResponse struct {
	BookId uuid.UUID `json:"book_id"`
}
