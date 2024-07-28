package http

import (
	"context"
	"github.com/labstack/gommon/log"
	"os"
	"os/signal"
	"syscall"
)

// NewContext - создать новый контекст приложения. Context - является аналогом CancellationToken
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
