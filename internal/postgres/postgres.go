package postgres

import (
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"

	"github.com/rmukhamet/core_test_task/internal/config"
)

func New(cfg *config.PG) (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx", cfg.URL)
	return db, err
}
