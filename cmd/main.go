package main

import (
	"go-template-microservice-v2/config"
	"go-template-microservice-v2/internal/configurations"
	"go-template-microservice-v2/internal/data/repositories"
	gormpg "go-template-microservice-v2/pkg/gorm_pg"
	"go-template-microservice-v2/pkg/http"
	echoserver "go-template-microservice-v2/pkg/http/server"
	"go-template-microservice-v2/server"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"

	. "go-template-microservice-v2/cmd/templates"

	. "go-template-microservice-v2/internal/features/get_order/endpoints"
	. "go-template-microservice-v2/internal/features/get_order/queries"
)

func main() {
	fx.New(
		fx.Options(
			fx.Provide(
				config.NewConfig,
				http.NewContext,
				gormpg.NewPgGorm,
				repositories.NewPgItemsRepository,
				repositories.NewPgPaymentRepository,
				repositories.NewPgOrderRepository,
				echoserver.NewEchoServer,
				validator.New,
				NewTemplateRenderer,
				NewGetOrderHandler,
				NewWebOrderHandler,
			),
			fx.Invoke(func(e *echo.Echo, handler *WebOrderHandler, renderer *TemplateRenderer) {
				e.Renderer = renderer

				e.POST("/order", handler.WebGetOrderHandler)
				e.GET("/find", handler.WebOrderFormHandler)
			}),
			fx.Invoke(configurations.ConfigEndpoints),
			fx.Invoke(server.RunServers),
			fx.Invoke(NewGetOrderHandler),
			fx.Invoke(NewWebOrderHandler),
		),
	).Run()
}
