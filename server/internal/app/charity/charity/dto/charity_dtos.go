package dto

import "github.com/neak-group/nikoogah/utils/uuid"

type RegisterCharityParams struct {
	Name          string
	Phone         string
	CityPhoneCode string
	Email         string
	Province      string
	City          string
	Address       string
	PostalCode    string
	NationalID    string
	EconomicID    string
	CEO           string
}

type ModifyCharityParams struct {
	ID            uuid.UUID
	Name          string
	Phone         string
	CityPhoneCode string
	Email         string
	Province      string
	City          string
	Address       string
	PostalCode    string
}

type FetchCharityParams struct {
	CharityID uuid.UUID
}

type AddRepresentativeParams struct {
	CharityID uuid.UUID
	UserID    uuid.UUID
}

type RemoveRepresentativeParams struct {
	CharityID uuid.UUID
	UserID    uuid.UUID
}

type CheckRepresentativeAccessParams struct {
	CharityID uuid.UUID
	UserID    uuid.UUID
	AccessKey string
}
