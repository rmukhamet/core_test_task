package storage

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/rmukhamet/core_test_task/internal/config"
	"github.com/rmukhamet/core_test_task/internal/controller"
	"github.com/rmukhamet/core_test_task/internal/grpcserver"
	"github.com/rmukhamet/core_test_task/internal/mq"
	pg "github.com/rmukhamet/core_test_task/internal/postgres"
	"github.com/rmukhamet/core_test_task/internal/repository"
)

type StorageControllerI interface {
	Save(ctx context.Context) error
}

type Storage struct {
	storageController StorageControllerI
	grpcServer        *grpcserver.GRPCService
}

func New(cfg *config.StorageConfig) *Storage {
	mq := mq.New(&cfg.REDIS)

	pgConn, err := pg.New(&cfg.PG)
	if err != nil {
		log.Fatal(err)
	}

	repository := repository.New(pgConn)
	sc := controller.NewStorageController(cfg, mq, repository)

	grpcServer := grpcserver.New(cfg, repository)

	return &Storage{
		storageController: sc,
		grpcServer:        grpcServer,
	}
}

func (s *Storage) Run(ctx context.Context) error {
	err := s.storageController.Save(ctx)
	if err != nil {
		return fmt.Errorf("error saving data with error: %w", err)
	}

	err = s.grpcServer.Run(ctx)
	if err != nil {
		return fmt.Errorf("error grpc run with error: %w", err)
	}

	return nil
}

func (s *Storage) Init(cfg *config.StorageConfig) error {
	databaseURL := cfg.PG.URL

	databaseURL += "?sslmode=disable"

	var (
		attempts = 2
		err      error
		m        *migrate.Migrate
	)

	for attempts > 0 {
		m, err = migrate.New("file://migrations", databaseURL)
		if err == nil {
			break
		}

		log.Printf("Migrate: postgres is trying to connect, attempts left: %d", attempts)
		time.Sleep(3 * time.Second)
		attempts--
	}

	if err != nil {
		log.Fatalf("Migrate: postgres connect error: %s", err)
	}

	err = m.Up()
	defer m.Close()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("Migrate: up error: %s", err)
	}

	if errors.Is(err, migrate.ErrNoChange) {
		log.Printf("Migrate: no change")
		return err
	}

	log.Printf("Migrate: up success")

	return nil
}

func (s *Storage) Close(ctx context.Context) error {
	err := s.grpcServer.Close(ctx)
	return err
}
