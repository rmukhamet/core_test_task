package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/rmukhamet/core_test_task/internal/config"
	"github.com/rmukhamet/core_test_task/internal/model"
)

type StorageController struct {
	mq   Subscriber
	repo RepositaryI
}

func NewStorageController(cfg *config.StorageConfig, mq Subscriber, repo RepositaryI) *StorageController {
	return &StorageController{
		mq:   mq,
		repo: repo,
	}
}

func (sc StorageController) Save(ctx context.Context) error {
	log.Print("subscribing")
	retailerCh, err := sc.mq.Subscribe(ctx)
	if err != nil {
		return fmt.Errorf("error with subscription: %w", err)
	}

	go func() {
		for task := range retailerCh {
			// todo check context done with select
			switch task.Type {
			case model.TaskTypeCreate:
				var retailer model.Retailer
				err := json.Unmarshal(task.Data.(json.RawMessage), &retailer)
				if err != nil {
					log.Printf("wrong data in task %+v, error: %s\n", task.Data, err.Error())
					continue
				}

				err = sc.repo.Create(ctx, retailer)
				if err != nil {
					log.Printf("failed create retailer %+v in db with error: %s\n", retailer, err.Error())
					continue
				}
			case model.TaskTypeUpdate:
				var retailer model.Retailer
				err := json.Unmarshal(task.Data.(json.RawMessage), &retailer)
				if err != nil {
					log.Printf("wrong data in task %+v, error: %s\n", task.Data, err.Error())
					continue
				}

				err = sc.repo.Update(ctx, retailer)
				if err != nil {
					log.Printf("failed update retailer %+v in db with error: %s\n", retailer, err.Error())
					continue
				}
			default:
				log.Printf("received unknown task type %+v\n", task)
				continue
			}
		}
	}()

	return nil
}
