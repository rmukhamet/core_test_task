package storage

import (
	"context"

	"github.com/rmukhamet/core_test_task/internal/config"
)

type Storage struct {
}

func New(cfg *config.StorageConfig) *Storage {
	return &Storage{}
}

func (s *Storage) Init() error {

	return nil
}

func (s *Storage) Run() error {

	return nil
}

func (s *Storage) Close(ctx context.Context) error {

	return nil
}
