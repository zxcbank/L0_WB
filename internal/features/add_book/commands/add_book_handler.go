package commands

import (
	"context"
	"go-template-microservice-v2/internal/data/contracts"
	"go-template-microservice-v2/internal/data/entities"
)

// AddBookHandler - хендлер для команды AddUserRequestCommand
type AddBookHandler struct {
	Repository contracts.IBookRepository
	Ctx        context.Context
}

// NewAddBookHandler - DI
func NewAddBookHandler(
	repository contracts.IBookRepository,
	ctx context.Context) *AddBookHandler {
	return &AddBookHandler{Repository: repository, Ctx: ctx}
}

// Handle - выполнить
func (handler *AddBookHandler) Handle(ctx context.Context, command *AddBookCommand) (*AddBookResponse, error) {
	bookEntity := entities.CreateBookEntity(
		command.Name,
		command.Author,
		command.Price)

	err := handler.Repository.AddBook(bookEntity)
	if err != nil {
		return nil, err
	}

	return &AddBookResponse{BookId: bookEntity.Id}, nil
}
