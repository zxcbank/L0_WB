package configurations

import (
	"context"
	addOrderEndpoints "go-template-microservice-v2/internal/features/add_order/endpoints"
	getOrderEndpoints "go-template-microservice-v2/internal/features/get_order/endpoints"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func ConfigEndpoints(validator *validator.Validate, echo *echo.Echo, ctx context.Context) {
	addOrderEndpoints.MapRoute(validator, echo, ctx)
	getOrderEndpoints.MapRoute(validator, echo, ctx)

}
