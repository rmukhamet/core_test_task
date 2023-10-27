package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rmukhamet/core_test_task/internal/gateway"

	"github.com/rmukhamet/core_test_task/internal/config"
)

func main() {
	err := run()
	if err != nil {
		log.Fatalf("app error: %s", err)
	}

	os.Exit(0)
}

func run() error {
	cfg, err := config.New()
	if err != nil {
		return fmt.Errorf("config error: %w", err)
	}

	app := gateway.New(cfg)

	// listen to OS signals and gracefully shutdown
	stopped := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-sigint
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := app.Close(ctx); err != nil {
			log.Printf("application shutdown error:: %v", err)
		}
		close(stopped)
	}()

	// err := app.Init()
	// if err != nil {
	// 	return fmt.Errorf("failed gateway init %w", err)
	// }
	err = app.Run()
	if err != nil {
		return fmt.Errorf("gateway error %w", err)
	}

	<-stopped

	log.Print("Application stopped")

	return nil
}
