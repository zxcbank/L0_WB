package http

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/gommon/log"
)

func NewContext() context.Context {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Info("context is canceled!")
				cancel()
				return
			}
		}
	}()
	return ctx
}
