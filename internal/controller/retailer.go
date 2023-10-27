package controller

import (
	"context"

	"github.com/rmukhamet/core_test_task/internal/apperrors"
	"github.com/rmukhamet/core_test_task/internal/config"
	"github.com/rmukhamet/core_test_task/internal/model"
)

type Publisher interface {
	Publish(ctx context.Context, v interface{}) error
}

type RetailerController struct {
	mq Publisher
}

func NewRetailerController(cfg *config.GatewayConfig, mq Publisher) *RetailerController {
	return &RetailerController{
		mq: mq,
	}
}

func (rc *RetailerController) Create(ctx context.Context, retailer model.Retailer) error {
	return rc.mq.Publish(ctx, retailer)
}

func (rc *RetailerController) Update(ctx context.Context, retailer model.Retailer) error {
	if retailer.Version.Version == 0 {
		return apperrors.ErrorWrongVersion
	}
	return nil
}
