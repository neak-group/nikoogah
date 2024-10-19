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

type CharityDTO struct {
	Name          string
	Phone         string
	Email         string
	Address       string
	PostalCode    string
	NationalID    string
	EconomicID    string
	CEO           string
}
