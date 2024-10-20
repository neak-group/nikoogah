package dto

import "github.com/neak-group/nikoogah/utils/uuid"

type RegisterCharityParams struct {
	Name          string `json:"name"`
	Phone         string `json:"phone"`
	CityPhoneCode string `json:"cityPhoneCode"`
	Email         string `json:"email"`
	Province      string `json:"province"`
	City          string `json:"city"`
	Address       string `json:"address"`
	PostalCode    string `json:"postalCode"`
	NationalID    string `json:"nationalId"`
	EconomicID    string `json:"economicId"`
	CEO           string `json:"ceo"`
}

type ModifyCharityParams struct {
	ID            uuid.UUID `json:"-"`
	Name          string    `json:"name"`
	Phone         string    `json:"phone"`
	CityPhoneCode string    `json:"cityPhoneCode"`
	Email         string    `json:"email"`
	Province      string    `json:"province"`
	City          string    `json:"city"`
	Address       string    `json:"address"`
	PostalCode    string    `json:"postalCode"`
}

type FetchCharityParams struct {
	CharityID uuid.UUID
}

type CharityDTO struct {
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Address    string `json:"address"`
	PostalCode string `json:"postalCode"`
	NationalID string `json:"nationalId"`
	EconomicID string `json:"economicId"`
	CEO        string `json:"ceo"`
}
