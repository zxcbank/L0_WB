package server

import (
	"context"
	"go-template-microservice-v2/config"
	echoserver "go-template-microservice-v2/pkg/http/server"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"go.uber.org/fx"
)

func RunServers(lc fx.Lifecycle, ctx context.Context, e *echo.Echo, cfg *config.Config) error {
	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			log.Println("Starting server")

			go func() {
				if err := echoserver.RunHttpServer(ctx, e, cfg.Echo); !errors.Is(err, http.ErrServerClosed) {
					log.Fatalf("error running http server: %v", err)
				}
			}()

			e.GET("/", func(c echo.Context) error {
				return c.String(http.StatusOK, cfg.ServiceName)
			})

			return nil
		},
		OnStop: func(_ context.Context) error {
			log.Println("all servers shutdown gracefully...")
			return nil
		},
	})

	return nil
}
