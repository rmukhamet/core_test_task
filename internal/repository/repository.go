package repository

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/rmukhamet/core_test_task/internal/model"
)

type Repository struct {
	conn *sqlx.DB
}

func New(conn *sqlx.DB) *Repository {
	return &Repository{
		conn: conn,
	}
}

func (r *Repository) Create(ctx context.Context, retailer model.Retailer) error {
	log.Print("repository create", retailer)
	result, err := r.conn.Exec(createRetailer,
		retailer.ID,
		retailer.Name,
		retailer.Address.City,
		retailer.Address.Street,
		retailer.Address.House,
		retailer.Owner.FirstName,
		retailer.Owner.LastName,
		retailer.OpenTime,
		retailer.CloseTime,
		retailer.Version.CreatedAt,
		retailer.Version.Actor,
	)
	if err != nil {
		return err
	}

	rowNum, err := result.RowsAffected()
	log.Print("created", rowNum)
	return err
}

func (r *Repository) Update(ctx context.Context, retailer model.Retailer) error {
	return nil //todo
}
func (r *Repository) GetRetailerList(ctx context.Context) ([]model.Retailer, error) {
	rows, err := r.conn.Queryx(listRetailer)
	if err != nil {
		return nil, nil
	}

	retailers := make(Retailers, 0, 0) // use limit and page and memory will be fine

	for rows.Next() {
		var retailer Retailer
		err = rows.StructScan(&retailer)
		if err != nil {
			return nil, err
		}
		retailers = append(retailers, retailer)
	}

	return retailers.ToDTO(), nil
}
func (r *Repository) GetRetailerByID(ctx context.Context, ID string) (model.Retailer, error) {
	// select id, max(modstamp) from test where modstamp <= <ref_time> and (del_modstamp is null || del_modstamp <= <ref_time>) group by id;

	return model.Retailer{
		ID: "sfsfsfsfsfsdfdsfsfs",
	}, nil
}

func (r *Repository) DeleteRetailer(ctx context.Context, ID string) error {
	return nil
}

func (r *Repository) DeleteRetailerVersion(ctx context.Context, ID string, version int) error {
	return nil
}

func (r *Repository) History(ctx context.Context, ID string) ([]model.Retailer, error) {
	return nil, nil
}

func (r *Repository) GetRetailerVersion(ctx context.Context, ID string, version int) (model.Retailer, error) {
	return model.Retailer{}, nil
}
