package gateway

import (
	"context"
	"fmt"

	"github.com/rmukhamet/core_test_task/internal/config"
	"github.com/rmukhamet/core_test_task/internal/controller"
	"github.com/rmukhamet/core_test_task/internal/mq"
	"github.com/rmukhamet/core_test_task/internal/webserver"
)

type Gateway struct {
	webserver *webserver.WebServer
}

func New(cfg *config.GatewayConfig) *Gateway {
	mq := mq.New(&cfg.REDIS)
	rc := controller.NewRetailerController(cfg, mq)

	ws := webserver.New(cfg, rc)
	return &Gateway{
		webserver: ws,
	}
}
func (gw *Gateway) Run() error {
	err := gw.webserver.Run()
	if err != nil {
		return fmt.Errorf("failed run webserver: %w", err)
	}

	return nil
}

func (gw *Gateway) Close(ctx context.Context) error {
	err := gw.webserver.Close(ctx)
	if err != nil {
		return fmt.Errorf("failed shutdown webserver: %w", err)
	}

	return nil
}
