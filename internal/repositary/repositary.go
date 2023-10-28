package repositary

import (
	"context"
	"database/sql"

	"github.com/rmukhamet/core_test_task/internal/model"
)

type Repositary struct {
	conn *sql.DB
}

func New(conn *sql.DB) *Repositary {
	return &Repositary{
		conn: conn,
	}
}

func (r *Repositary) Create(ctx context.Context, retailer model.Retailer) error {
	return nil
}

func (r *Repositary) Update(ctx context.Context, retailer model.Retailer) error {
	return nil
}
