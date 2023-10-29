package webserver

import (
	"context"

	"github.com/rmukhamet/core_test_task/internal/model"
)

type RetailerControllerI interface {
	RetailerCreator
	RetailerUpdator
	RetailerGetter
}

type RetailerCreator interface {
	Create(ctx context.Context, retailer model.Retailer) error
}

type RetailerUpdator interface {
	Update(ctx context.Context, retailer model.Retailer) error
}

type RetailerGetter interface {
	GetRetailerByID(ctx context.Context, ID string) (model.Retailer, error)
}

type VersionLister interface {
	ListVersion(ctx context.Context, ID string) ([]model.Version, error)
}

type VersionGetter interface {
	GetVersion(ctx context.Context, ID string, version int) (model.Version, error)
}
