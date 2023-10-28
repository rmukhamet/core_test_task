package storage

import (
	"context"
	"fmt"

	"github.com/rmukhamet/core_test_task/internal/config"
	"github.com/rmukhamet/core_test_task/internal/controller"
	"github.com/rmukhamet/core_test_task/internal/mq"
	"github.com/rmukhamet/core_test_task/internal/postgres"
	"github.com/rmukhamet/core_test_task/internal/repositary"
)

type StorageControllerI interface {
	Save(ctx context.Context) error
}

type Storage struct {
	storageController StorageControllerI
}

func New(cfg *config.StorageConfig) *Storage {
	mq := mq.New(&cfg.REDIS)
	pgConn := postgres.New(&cfg.PG)
	repositary := repositary.New(pgConn)
	sc := controller.NewStorageController(cfg, mq, repositary)
	return &Storage{
		storageController: sc,
	}
}

func (s *Storage) Run(ctx context.Context) error {
	err := s.storageController.Save(ctx)
	if err != nil {
		return fmt.Errorf("error saving data with error: %w", err)
	}

	return nil
}

func (s *Storage) Init() error {
	//db migrate
	return nil
}

func (s *Storage) Close(ctx context.Context) error {

	return nil
}
