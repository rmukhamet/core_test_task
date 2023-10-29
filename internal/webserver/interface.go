package webserver

import (
	"context"

	"github.com/rmukhamet/core_test_task/internal/model"
)

type RetailerControllerI interface {
	RetailerCreatorUpdator
	RetailerGetter
	RetailerDeleter
}

type RetailerCreatorUpdator interface {
	Create(ctx context.Context, retailer model.Retailer) error
	Update(ctx context.Context, retailer model.Retailer) error
}

type RetailerGetter interface {
	GetRetailerByID(ctx context.Context, ID string) (model.Retailer, error)
	GetRetailerList(ctx context.Context) ([]model.Retailer, error)
	GetRetailerVersionList(ctx context.Context, retailerID string) ([]model.Version, error)
	GetRetailerVersion(ctx context.Context, retailerID string, versionID int) (model.Retailer, error)
}

type RetailerDeleter interface {
	DeleteRetailerVersion(ctx context.Context, retailerID string, versionID int) error
	DeleteRetailer(ctx context.Context, retailerID string) error
}
