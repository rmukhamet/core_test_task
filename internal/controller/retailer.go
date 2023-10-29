package controller

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/rmukhamet/core_test_task/internal/apperrors"
	"github.com/rmukhamet/core_test_task/internal/config"
	"github.com/rmukhamet/core_test_task/internal/model"
)

type RetailerController struct {
	mq Publisher
	tr TransportI
}

func NewRetailerController(cfg *config.GatewayConfig, mq Publisher, tr TransportI) *RetailerController {
	return &RetailerController{
		mq: mq,
		tr: tr,
	}
}

func (rc *RetailerController) Create(ctx context.Context, retailer model.Retailer) error {
	retailer.ID = uuid.New().String()
	retailer.Version.CreatedAt = time.Now()

	task := model.NewTask(model.TaskTypeCreate, retailer)
	return rc.mq.Publish(ctx, task)
}

func (rc *RetailerController) Update(ctx context.Context, retailer model.Retailer) error {
	if retailer.Version.Version == 0 {
		return apperrors.ErrorWrongVersion
	}

	if retailer.ID == "" {
		return apperrors.ErrorRetailerIDRequired
	}

	task := model.NewTask(model.TaskTypeUpdate, retailer)
	return rc.mq.Publish(ctx, task)
}

func (rc *RetailerController) GetRetailerByID(ctx context.Context, retailerID string) (model.Retailer, error) {

	return model.Retailer{}, nil
}

func (rc *RetailerController) GetRetailerList(ctx context.Context) ([]model.Retailer, error) {
	return nil, nil
}

func (rc *RetailerController) DeleteRetailer(ctx context.Context, retailerID string) error {
	return nil
}

func (rc *RetailerController) GetRetailerVersionList(ctx context.Context, retailerID string) ([]model.Version, error) {
	return nil, nil
}

func (rc *RetailerController) GetRetailerVersion(ctx context.Context, retailerID string, versionID int) (model.Retailer, error) {
	return model.Retailer{}, nil
}

func (rc *RetailerController) DeleteRetailerVersion(ctx context.Context, retailerID string, versionID int) error {
	return nil
}
