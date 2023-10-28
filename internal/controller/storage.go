package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/rmukhamet/core_test_task/internal/config"
)

type StorageController struct {
	mq   Subscriber
	repo RepositaryI
}

func NewStorageController(cfg *config.StorageConfig, mq Subscriber, repo RepositaryI) *StorageController {
	return &StorageController{
		mq: mq,
	}
}

func (sc StorageController) Save(ctx context.Context) error {
	log.Print("subscribing")
	retailerCh, err := sc.mq.Subscribe(ctx, "ddd")
	if err != nil {
		return fmt.Errorf("error with subscription: %w", err)
	}

	for retailer := range retailerCh {
		err := sc.repo.Create(ctx, retailer)
		if err != nil {
			log.Print("failed create in db", err)
			continue
		}
	}

	return nil
}
