package endpoints

import (
	"context"
	"go-template-microservice-v2/internal/features/get_order/queries"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/mehdihadeli/go-mediatr"
)

// MapRoute - настройка маршрутизации
func MapRoute(validator *validator.Validate, echo *echo.Echo, ctx context.Context) {
	group := echo.Group("/api/v1/order")
	group.GET("/:id", GetOrderById(validator, ctx))
}

// AddBook
// @Tags        Order
// @Summary     Get Order
// @Description Get Order from catalogue
// @Accept      json
// @Produce     json
// @Param       GetOrderQuery body queries.GetOrderQuery true "Book data"
// @Success     200  {object} queries.GetOrderResponse
// @Security -
// @Router      /api/v1/Order/{id} [get]
func GetOrderById(validator *validator.Validate, ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		query := queries.GetOrderQuery{Id: uuid.MustParse(id)}

		result, err := mediatr.Send[*queries.GetOrderQuery, *queries.GetOrderResponse](ctx, &query)

		if err != nil {
			log.Errorf("(Handle) err: {%v}", err)
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}
		// для html

		//err = c.Render(http.StatusOK, "order_info.html", map[string]interface{}{
		//	"Order": result,
		//})
		//if err != nil {
		//	return err
		//}

		return c.JSON(http.StatusCreated, result)
	}
}
