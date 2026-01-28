package domain

import (
	"time"
)

type Customer struct {
	ID          int64
	Firstname   string
	Lastname    string
	Genda       Genda
	Email       string
	PhoneNumber string
	Status      Status
	AddressID   int64
	CreatedAt   time.Time  //ex:2026-01-07 14:42:18.123456789 +0000 UTC
	UpdatedAt   *time.Time //because it is not mandatory
}

type Genda string

const (
	Female Genda = "F"
	Male   Genda = "M"
)

type Status string

const (
	Active    Status = "ACTIVE"
	Suspended Status = "SUSPENDED"
	Deleted   Status = "DELETED"
)

// PatchBusinessCustomer pour la mise à jour de certains champs
type PatchBusinessCustomer struct {
	Firstname   *string
	Lastname    *string
	Email       *string
	PhoneNumber *string
	AddressID   *int64
}

func (bs *Customer) ApplyPatchCustomer(patch PatchBusinessCustomer) {
	if patch.Firstname != nil {
		bs.Firstname = *patch.Firstname
	}
	if patch.Lastname != nil {
		bs.Lastname = *patch.Lastname
	}
	if patch.Email != nil {
		bs.Email = *patch.Email
	}
	if patch.PhoneNumber != nil {
		bs.PhoneNumber = *patch.PhoneNumber
	}
	if patch.AddressID != nil {
		bs.AddressID = *patch.AddressID
	}

	bs.applyPatchUpdateDate()
}

func (bs *Customer) applyPatchUpdateDate() {
	now := time.Now()
	bs.UpdatedAt = &now
}
