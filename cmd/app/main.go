package main

import (
	"context"
	"errors"
	"github.com/Alp4ka/mlogger"
	"github.com/Alp4ka/mlogger/field"
	_ "github.com/jackc/pgx/v5/stdlib"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	_defaultGracefulShutdownTimeout = 15 * time.Second
)

func main() {
	env := setup()
	go awaitGracefulShutdown(env.ctx, env.cancelFunc)

	mlogger.L().Info("Starting app")
	err := env.app.Run(env.ctx)
	mlogger.L().Info("App stopped!", field.Error(err))

	err = env.app.Close()
	if err != nil {
		mlogger.L().Error("Failed to close app", field.Error(err))
	}
	mlogger.L().Info("App closed successfully!")
}

func awaitGracefulShutdown(ctx context.Context, cancel context.CancelFunc) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer signal.Stop(ch)

	select {
	case sig := <-ch:
		mlogger.L().Info("graceful shutdown: reason: received signal", field.Any("signal", sig.String()))
	}
	cancel()

	go func() {
		time.Sleep(_defaultGracefulShutdownTimeout)
		mlogger.L().Fatal("force quit, graceful shutdown timeout expired, force exit", field.String("timeout", _defaultGracefulShutdownTimeout.String()))
	}()

	cause := context.Cause(ctx)
	if !errors.Is(cause, context.Canceled) {
		mlogger.L().Error("app finished with fail, cause: %s", field.Error(cause))
		return
	}
	mlogger.L().Info("app finished successfully!")
}
