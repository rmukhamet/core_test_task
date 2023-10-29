package controller

import (
	"context"

	"github.com/rmukhamet/core_test_task/internal/apperrors"
	"github.com/rmukhamet/core_test_task/internal/config"
	"github.com/rmukhamet/core_test_task/internal/model"
)

type RetailerController struct {
	mq Publisher
}

func NewRetailerController(cfg *config.GatewayConfig, mq Publisher) *RetailerController {
	return &RetailerController{
		mq: mq,
	}
}

func (rc *RetailerController) Create(ctx context.Context, retailer model.Retailer) error {
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
