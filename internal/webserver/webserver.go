package webserver

import (
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/rmukhamet/core_test_task/internal/config"
)

type WebServer struct {
	server             *fiber.App
	port               string
	retailerController RetailerControllerI
}

func New(cfg *config.GatewayConfig, retailerController RetailerControllerI) *WebServer {
	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "GW",
		AppName:       "gateway-service",
	})

	ws := &WebServer{
		server:             app,
		port:               cfg.HTTP.Port,
		retailerController: retailerController,
	}

	ws.router()

	return ws
}

func (ws *WebServer) Run() error {
	err := ws.server.Listen(fmt.Sprintf(":%s", ws.port))
	if err != nil {
		log.Fatalf("HTTP server Error: %v", err)
	}

	return err
}

func (ws *WebServer) Close(ctx context.Context) error {
	return ws.server.ShutdownWithContext(ctx)
}
