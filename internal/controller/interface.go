package controller

import (
	"context"

	"github.com/rmukhamet/core_test_task/internal/model"
)

type Publisher interface {
	Publish(ctx context.Context, t model.Task) error
}

type Subscriber interface {
	Subscribe(ctx context.Context) (<-chan model.Task, error)
}

type RepositaryI interface {
	Create(ctx context.Context, retailer model.Retailer) error
	Update(ctx context.Context, retailer model.Retailer) error
}

type TransportI interface {
	GetRetailerByID(ctx context.Context, ID string) (model.Retailer, error)
	GetRetailerList(ctx context.Context) ([]model.Retailer, error)
	GetRetailerVersionList(ctx context.Context, retailerID string) ([]model.Retailer, error)
	GetRetailerVersion(ctx context.Context, retailerID string, versionID int) (model.Retailer, error)
	DeleteRetailerVersion(ctx context.Context, retailerID string, versionID int) error
	DeleteRetailer(ctx context.Context, retailerID string) error
}
