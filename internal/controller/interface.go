package controller

import (
	"context"

	"github.com/rmukhamet/core_test_task/internal/model"
)

type Publisher interface {
	Publish(ctx context.Context, v interface{}) error
}

type Subscriber interface {
	Subscribe(ctx context.Context, v interface{}) (<-chan model.Retailer, error)
}

type RepositaryI interface {
	Create(ctx context.Context, retailer model.Retailer) error
	Update(ctx context.Context, retailer model.Retailer) error
}
