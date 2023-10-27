package webserver

import "time"

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
}

type RetailerUpdateRequest struct {
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
