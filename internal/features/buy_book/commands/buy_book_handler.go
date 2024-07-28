package commands

import (
	"context"
	"go-template-microservice-v2/internal/data/contracts"
)

// BuyBookHandler - хендлер для команды AddUserRequestCommand
type BuyBookHandler struct {
	Repository contracts.IBookRepository
	Ctx        context.Context
}

// NewBuyBookHandler - DI
func NewBuyBookHandler(
	repository contracts.IBookRepository,
	ctx context.Context) *BuyBookHandler {
	return &BuyBookHandler{Repository: repository, Ctx: ctx}
}

// Handle - выполнить
func (handler *BuyBookHandler) Handle(ctx context.Context, command *BuyBookCommand) (*BuyBookResponse, error) {
	book, err := handler.Repository.GetBook(command.BookId)

	if err != nil {
		return nil, err
	}

	book.Enabled = false

	err = handler.Repository.UpdateBook(*book)
	if err != nil {
		return nil, err
	}

	return &BuyBookResponse{Result: book.Enabled}, nil
}
