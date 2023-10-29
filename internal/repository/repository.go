package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/rmukhamet/core_test_task/internal/model"
)

type Repository struct {
	conn *sql.DB
}

func New(conn *sql.DB) *Repository {
	return &Repository{
		conn: conn,
	}
}

func (r *Repository) Create(ctx context.Context, retailer model.Retailer) error {
	log.Print("repository create", retailer)
	return nil //todo
}

func (r *Repository) Update(ctx context.Context, retailer model.Retailer) error {
	return nil //todo
}

func (r *Repository) GetRetailerByID(ctx context.Context, ID string) (model.Retailer, error) {
	//select id, max(modstamp) from test where modstamp <= <ref_time> and (del_modstamp is null || del_modstamp <= <ref_time>) group by id;

	return model.Retailer{}, nil
}

func (r *Repository) DeleteRetailer(ctx context.Context, ID string) error {
	return nil
}

func (r *Repository) DeleteVersion(ctx context.Context, ID string, version int) error {
	return nil
}

func (r *Repository) History(ctx context.Context, ID string) ([]model.Retailer, error) {
	return nil, nil
}

func (r *Repository) GetRetailerVersion(ctx context.Context, ID string, version int) (model.Retailer, error) {
	return model.Retailer{}, nil
}
