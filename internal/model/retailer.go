package model

import "time"

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
	Creator string
	Version int
}
