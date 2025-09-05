package main

import (
	"go-template-microservice-v2/config"
	"go-template-microservice-v2/internal/data/repositories"
	. "go-template-microservice-v2/internal/features/endpoints"
	. "go-template-microservice-v2/internal/features/queries"
	gormpg "go-template-microservice-v2/pkg/gorm_pg"
	"go-template-microservice-v2/pkg/http"
	echoserver "go-template-microservice-v2/pkg/http/server"
	"go-template-microservice-v2/server"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"

	. "go-template-microservice-v2/cmd/templates"
)

func main() {
	fx.New(
		fx.Options(
			fx.Provide(
				config.NewConfig,
				http.NewContext,
				gormpg.NewPgGorm,
				repositories.NewPgOrderRepository,
				echoserver.NewEchoServer,
				validator.New,
				NewTemplateRenderer,
				NewOrderService,
				NewOrderEndpoint,
			),
			fx.Invoke(func(e *echo.Echo, handler *OrderEndpoint, renderer *TemplateRenderer) {
				e.Renderer = renderer

				e.POST("/order", handler.OrderShowResult)
				e.GET("/find", handler.OrderForm)
			}),
			fx.Invoke(server.RunServers),
			fx.Invoke(NewOrderService),
			fx.Invoke(NewOrderEndpoint),
		),
	).Run()
}
