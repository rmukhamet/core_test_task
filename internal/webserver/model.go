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

type RetailerUpdateRequest struct { //todo json tags
	ID             string
	OwnerFirstName string
	ownerLastName  string
	OpenTime       time.Time
	CloseTime      time.Time

	Actor   string
	Version int
}

func (rur *RetailerUpdateRequest) ToDTO() model.Retailer {
	return model.Retailer{
		ID: rur.ID,
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

type RetailerResponse struct { //todo json tags
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
type Version struct { //todo json tags
	Version   int
	Actor     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewRetailerResponse(r model.Retailer) RetailerResponse {
	return RetailerResponse{
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

type RetailerVersionListResponse struct {
	Items []Version `json:"items"`
}

func NewRetalerVersionListResponse(versions []model.Version) RetailerVersionListResponse {
	versionsResp := make([]Version, 0, len(versions))

	for _, v := range versions {
		versionsResp = append(versionsResp,
			Version{
				Version:   v.Version,
				Actor:     v.Actor,
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
			})
	}

	return RetailerVersionListResponse{
		Items: versionsResp,
	}
}

type RetailerListResponse struct {
	Items []RetailerResponse `json:"items"`
}

func NewRetalerListResponse(retailers []model.Retailer) RetailerListResponse {
	retailersResp := make([]RetailerResponse, 0, len(retailers))

	for _, r := range retailers {
		retailersResp = append(retailersResp, NewRetailerResponse(r))
	}

	return RetailerListResponse{
		Items: retailersResp,
	}
}
