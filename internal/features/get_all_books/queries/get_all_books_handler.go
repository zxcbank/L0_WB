package queries

import (
	"context"
	"go-template-microservice-v2/internal/data/contracts"
)

type GetAllBooksHandler struct {
	Repository contracts.IBookRepository
	Ctx        context.Context
}

// NewGetAllBooksHandler - DI
func NewGetAllBooksHandler(
	repository contracts.IBookRepository,
	ctx context.Context) *GetAllBooksHandler {
	return &GetAllBooksHandler{Repository: repository, Ctx: ctx}
}

// Handle - выполнить
func (handler *GetAllBooksHandler) Handle(ctx context.Context, command *GetAllBooksQuery) (*GetAllBooksResponse, error) {
	getAllBooksResponse := &GetAllBooksResponse{
		Books: make([]GetAllBooksResponseItem, 0),
	}

	result, err := handler.Repository.GetAllBook()
	if err != nil {
		return nil, err
	}

	for _, element := range result {
		getAllBooksResponse.Books = append(getAllBooksResponse.Books, GetAllBooksResponseItem{
			Id:      element.Id,
			Name:    element.Name,
			Author:  element.Author,
			Price:   element.Price,
			Enabled: element.Enabled,
		})
	}

	return getAllBooksResponse, nil
}
