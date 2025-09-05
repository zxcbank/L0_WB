package echoserver

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

const (
	MaxHeaderBytes = 1 << 20
	ReadTimeout    = 15 * time.Second
	WriteTimeout   = 15 * time.Second
)

type EchoConfig struct {
	Port                string   `mapstructure:"port" validate:"required"`
	Development         bool     `mapstructure:"development"`
	BasePath            string   `mapstructure:"basePath" validate:"required"`
	DebugErrorsResponse bool     `mapstructure:"debugErrorsResponse"`
	IgnoreLogUrls       []string `mapstructure:"ignoreLogUrls"`
	Timeout             int      `mapstructure:"timeout"`
	Host                string   `mapstructure:"host"`
}

func NewEchoServer() *echo.Echo {
	e := echo.New()
	return e
}

func RunHttpServer(ctx context.Context, echo *echo.Echo, cfg *EchoConfig) error {
	echo.Server.ReadTimeout = ReadTimeout
	echo.Server.WriteTimeout = WriteTimeout
	echo.Server.MaxHeaderBytes = MaxHeaderBytes

	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Infof("Server is shutting down. HTTP POST: {%s}", cfg.Port)
				err := echo.Shutdown(ctx)
				if err != nil {
					log.Errorf("(Shutdown) err: {%v}", err)
					return
				}
				log.Info("server exited properly")
				return
			}
		}
	}()

	err := echo.Start(cfg.Port)

	return err
}
