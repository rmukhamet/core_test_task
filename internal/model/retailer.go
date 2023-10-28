package model

import "time"

type (
	Retailer struct {
		// constant fields
		ID      string
		Name    string
		Address Address
		// keep version
		Owner     Person
		OpenTime  time.Time
		CloseTime time.Time
		Version   Version
	}

	Address struct {
		City   string
		Street string
		House  string
	}

	Person struct {
		FirstName string
		LastName  string
	}

	Version struct {
		Actor   string
		Version int
	}
)

func (r Retailer) Validate() error {
	// TODO
	return nil
}
