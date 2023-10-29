package webserver

import (
	"time"

	"github.com/rmukhamet/core_test_task/internal/model"
)

type RetailerCreateRequest struct {
	Name          string `json:"name"`
	AddressCity   string `json:"address_city"`
	AddressStreet string `json:"address_street"`
	AddressHouse  string `json:"address_house"`
	// keep version
	OwnerFirstName string    `json:"owner_first_name"`
	OwnerLastName  string    `json:"owner_last_name"`
	OpenTime       time.Time `json:"open_time"`
	CloseTime      time.Time `json:"close_time"`

	Actor string `json:"-"`
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
			LastName:  rcr.OwnerLastName,
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

type RetailerGetResponse struct {
	ID            string
	Name          string
	AddressCity   string
	AddressStreet string
	AddressHouse  string
	// keep version
	OwnerFirstName string
	ownerLastName  string
	OpenTime       time.Time
	CloseTime      time.Time

	Version   int
	Actor     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewRetailerGetResponse(r model.Retailer) RetailerGetResponse {
	return RetailerGetResponse{
		ID:            r.ID,
		Name:          r.Name,
		AddressCity:   r.Address.City,
		AddressStreet: r.Address.Street,
		AddressHouse:  r.Address.House,
		// keep version
		OwnerFirstName: r.Owner.FirstName,
		ownerLastName:  r.Owner.LastName,
		OpenTime:       r.OpenTime,
		CloseTime:      r.CloseTime,

		Version:   r.Version.Version,
		Actor:     r.Version.Actor,
		CreatedAt: r.Version.CreatedAt,
		UpdatedAt: r.Version.UpdatedAt,
	}
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
