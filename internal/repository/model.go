package repository

import "github.com/rmukhamet/core_test_task/internal/model"

type Retailer struct {
	ID string `db:"id"`
}

func (r Retailer) ToDTO() model.Retailer {
	return model.Retailer{
		ID: r.ID,
	}
}

type Retailers []Retailer

func (rs Retailers) ToDTO() []model.Retailer {
	retailers := make([]model.Retailer, 0, len(rs))

	for _, r := range rs {
		retailers = append(retailers, r.ToDTO())
	}
	return retailers
}
