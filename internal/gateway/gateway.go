package gateway

import (
	"context"

	"github.com/rmukhamet/core_test_task/internal/config"
	"github.com/rmukhamet/core_test_task/internal/model"
	"github.com/rmukhamet/core_test_task/internal/webserver"
)

type Gateway struct {
	webserver *webserver.WebServer
}

func New(cfg *config.GatewayConfig) *Gateway {
	ws := webserver.New(cfg)
	return &Gateway{
		webserver: ws,
	}
}
func (gw *Gateway) Run() error {
	err := gw.webserver.Run()
	if err != nil {
		return err
	}

	return nil
}

func (gw *Gateway) Close(ctx context.Context) error {
	err := gw.webserver.Close(ctx)
	if err != nil {
		return err
	}

	return nil
}

type RetailerCreator interface {
	Create(ctx context.Context, retailer model.Retailer) error
}

type RetailerUpdator interface {
	Update(ctx context.Context, retailer model.Retailer) error
}

type RetailerGetter interface {
	GetRetail(ctx context.Context, ID string) (model.Retailer, error)
}

type VersionLister interface {
	ListVersion(ctx context.Context, ID string) ([]model.Version, error)
}

type VersionGetter interface {
	GetVersion(ctx context.Context, ID string, version int) (model.Version, error)
}
