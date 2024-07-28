package endpoints

import (
	"context"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/mehdihadeli/go-mediatr"
	"go-template-microservice-v2/internal/features/get_all_books/queries"
	"net/http"
)

// MapRoute - настройка маршрутизации
func MapRoute(validator *validator.Validate, echo *echo.Echo, ctx context.Context) {
	group := echo.Group("/api/v1/books")
	group.GET("", getAllBooks(validator, ctx))
}

// AddBook
// @Tags        Book
// @Summary     Get All Books
// @Description Get All Books from catalogue
// @Accept      json
// @Produce     json
// @Param       GetAllBooksQuery body queries.GetAllBooksQuery true "Book data"
// @Success     200  {object} queries.GetAllBooksResponse
// @Security -
// @Router      /api/v1/books [get]
func getAllBooks(validator *validator.Validate, ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		query := queries.GetAllBooksQuery{}

		result, err := mediatr.Send[*queries.GetAllBooksQuery, *queries.GetAllBooksResponse](ctx, &query)

		if err != nil {
			log.Errorf("(Handle) err: {%v}", err)
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusCreated, result)
	}
}
