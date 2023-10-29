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
	GetRetailerByID(ctx context.Context, retailerID string) (model.Retailer, error)
}
