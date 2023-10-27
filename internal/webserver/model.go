package webserver

import (
	"time"

	"github.com/rmukhamet/core_test_task/internal/model"
)

type RetailerCreateRequest struct {
	Name          string
	AddressCity   string
	AddressStreet string
	AddressHouse  string
	// keep version
	OwnerFirstName string
	ownerLastName  string
	OpenTime       time.Time
	CloseTime      time.Time

	Actor string
}

func (rcr *RetailerCreateRequest) ToDTO() model.Retailer {
	return model.Retailer{
		Name: rcr.Name,
		Address: model.Address{
			City:   rcr.AddressCity,
			Street: rcr.AddressStreet,
			House:  rcr.AddressHouse,
		},
		Owner: model.Person{
			FirstName: rcr.OwnerFirstName,
			LastName:  rcr.ownerLastName,
		},
		OpenTime:  rcr.OpenTime,
		CloseTime: rcr.CloseTime,
		Version: model.Version{
			Actor: rcr.Actor,
		},
	}
}

type RetailerUpdateRequest struct {
	OwnerFirstName string
	ownerLastName  string
	OpenTime       time.Time
	CloseTime      time.Time

	Actor   string
	Version int
}

func (rur *RetailerUpdateRequest) ToDTO() model.Retailer {
	return model.Retailer{
		Owner: model.Person{
			FirstName: rur.OwnerFirstName,
			LastName:  rur.ownerLastName,
		},
		OpenTime:  rur.OpenTime,
		CloseTime: rur.CloseTime,
		Version: model.Version{
			Version: rur.Version,
			Actor:   rur.Actor,
		},
	}
}

type RetailerGetRequest struct {
	ID string
}

type RetailerGetResponse struct {
	Name          string
	AddressCity   string
	AddressStreet string
	AddressHouse  string
	// keep version
	OwnerFirstName string
	ownerLastName  string
	OpenTime       time.Time
	CloseTime      time.Time

	Version int
}

type RetailerGetVersionListRequest struct {
	ID string
}
type RetailerGetVersionListResponse struct {
	Versions []Version
}

type Version struct {
	Creator string
	Version int
}
