package postgres

import (
	"database/sql"

	"github.com/rmukhamet/core_test_task/internal/config"
)

func New(cfg *config.PG) *sql.DB {
	//sqlx.New
	return &sql.DB{}
}
