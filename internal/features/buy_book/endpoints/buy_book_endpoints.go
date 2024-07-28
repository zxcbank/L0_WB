package endpoints

import (
	"context"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/mehdihadeli/go-mediatr"
	"github.com/pkg/errors"
	"go-template-microservice-v2/internal/features/buy_book/commands"
	"net/http"
)

// MapRoute - настройка маршрутизации
func MapRoute(validator *validator.Validate, echo *echo.Echo, ctx context.Context) {
	group := echo.Group("/api/v1/books/buy")
	group.POST("", buyBook(validator, ctx))
}

// AddBook
// @Tags        Book
// @Summary     Buy Book
// @Description Buy Book in catalogue
// @Accept      json
// @Produce     json
// @Param       BuyBookCommand body commands.BuyBookCommand true "Book data"
// @Success     200  {object} commands.BuyBookResponse
// @Security -
// @Router      /api/v1/books/buy [post]
func buyBook(validator *validator.Validate, ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		request := &commands.BuyBookCommand{}

		if err := c.Bind(request); err != nil {
			badRequestErr := errors.Wrap(err, "[addBookEndpoint_handler.Bind] error in the binding request")
			log.Error(badRequestErr)
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		if err := validator.StructCtx(ctx, request); err != nil {
			validationErr := errors.Wrap(err, "[addBook_handler.StructCtx] command validation failed")
			log.Error(validationErr)
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		result, err := mediatr.Send[*commands.BuyBookCommand, *commands.BuyBookResponse](ctx, request)

		if err != nil {
			log.Errorf("(Handle) err: {%v}", err)
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		log.Infof("(auto added) id: {%s}", result.Result)
		return c.JSON(http.StatusCreated, result)
	}
}
