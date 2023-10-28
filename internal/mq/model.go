package mq

import (
	"encoding/json"
	"time"

	"github.com/rmukhamet/core_test_task/internal/model"
)

type Retailer struct {
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

func (r Retailer) Validate() error {
	// TODO
	return nil
}

type Address struct {
	City   string
	Street string
	House  string
}

type Person struct {
	FirstName string
	LastName  string
}

type Version struct {
	Actor   string
	Version int
}

func NewRetailer(r model.Retailer) Retailer {
	return Retailer{
		Name:      r.Name,
		Address:   Address(r.Address),
		Owner:     Person(r.Owner),
		OpenTime:  r.OpenTime,
		CloseTime: r.CloseTime,
		Version: Version{
			Actor: r.Version.Actor,
		},
	}
}
func (r Retailer) MarshalBinary() (data []byte, err error) {
	bytes, err := json.Marshal(r)
	return bytes, err
}

func (r Retailer) ToDTO() model.Retailer {
	return model.Retailer{
		Name: r.Name,
		Address: model.Address{
			City:   r.Address.City,
			Street: r.Address.Street,
			House:  r.Address.House,
		},
		Owner: model.Person{
			FirstName: r.Owner.FirstName,
			LastName:  r.Owner.LastName,
		},
		OpenTime:  r.OpenTime,
		CloseTime: r.CloseTime,
		Version: model.Version{
			Actor:   r.Version.Actor,
			Version: r.Version.Version,
		},
	}

}
