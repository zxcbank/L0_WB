package configurations

import (
	"context"
	"github.com/mehdihadeli/go-mediatr"
	"go-template-microservice-v2/internal/data/contracts"
	addBookCommand "go-template-microservice-v2/internal/features/add_book/commands"
	buyBookCommand "go-template-microservice-v2/internal/features/buy_book/commands"
	getAllBooksQueries "go-template-microservice-v2/internal/features/get_all_books/queries"
)

// ConfigMediator - DI
func ConfigMediator(
	ctx context.Context,
	repository contracts.IBookRepository) (err error) {

	err = mediatr.RegisterRequestHandler[
		*addBookCommand.AddBookCommand,
		*addBookCommand.AddBookResponse](addBookCommand.NewAddBookHandler(repository, ctx))

	err = mediatr.RegisterRequestHandler[
		*buyBookCommand.BuyBookCommand,
		*buyBookCommand.BuyBookResponse](buyBookCommand.NewBuyBookHandler(repository, ctx))

	err = mediatr.RegisterRequestHandler[
		*getAllBooksQueries.GetAllBooksQuery,
		*getAllBooksQueries.GetAllBooksResponse](getAllBooksQueries.NewGetAllBooksHandler(repository, ctx))

	if err != nil {
		return err
	}

	return nil
}
