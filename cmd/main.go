package main

import (
	"github.com/go-playground/validator"
	"go-template-microservice-v2/config"
	"go-template-microservice-v2/internal/configurations"
	"go-template-microservice-v2/internal/data/entities"
	"go-template-microservice-v2/internal/data/repositories"
	gormpg "go-template-microservice-v2/pkg/gorm_pg"
	"go-template-microservice-v2/pkg/http"
	echoserver "go-template-microservice-v2/pkg/http/server"
	"go-template-microservice-v2/server"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Options(
			fx.Provide(
				config.NewConfig,
				http.NewContext,
				gormpg.NewPgGorm,
				repositories.NewPgBookRepository,
				echoserver.NewEchoServer,
				validator.New,
			),
			fx.Invoke(configurations.ConfigEndpoints),
			fx.Invoke(configurations.ConfigMediator),
			fx.Invoke(server.RunServers),
			fx.Invoke(
				func(sql *gormpg.PgGorm) error {
					return gormpg.Migrate(sql.DB, &entities.BookEntity{})
				}),
		),
	).Run()
}
