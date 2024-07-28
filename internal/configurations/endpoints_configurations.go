package configurations

import (
	"context"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	addBookEndpoints "go-template-microservice-v2/internal/features/add_book/endpoints"
	buyBookEndpoints "go-template-microservice-v2/internal/features/buy_book/endpoints"
	getAllBooksEndpoints "go-template-microservice-v2/internal/features/get_all_books/endpoints"
)

// ConfigEndpoints - конфигурирование ендпоинтов нашего API
func ConfigEndpoints(validator *validator.Validate, echo *echo.Echo, ctx context.Context) {
	addBookEndpoints.MapRoute(validator, echo, ctx)
	buyBookEndpoints.MapRoute(validator, echo, ctx)
	getAllBooksEndpoints.MapRoute(validator, echo, ctx)
}
