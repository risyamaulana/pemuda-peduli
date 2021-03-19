package infrastructure

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// CaptureSignal ...
func CaptureSignal() context.Context {
	// ctx, cancel := context.WithTimeout(context.Background(), time.Duration(30)*time.Second)
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

		select {
		case <-sigs:
			cancel()
		case <-ctx.Done():
			return
		}

		defer log.Println("🔴 Termination request in 10 secs. Press ctrl+c again to force termination.")
	}()

	return ctx
}

// GracefulShutdown ...
func GracefulShutdown() context.Context {
	// Give 30 second timeout before forcing shutdown
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)

	go func() {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

		select {
		case <-sigs:
			cancel()
		case <-ctx.Done():
			return
		}
	}()

	return ctx
}
