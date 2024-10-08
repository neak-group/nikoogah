package model

import (
	"github.com/neak-group/nikoogah/internal/app/user/entity"
)

type UserInput struct {
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	PhoneNumber  string `json:"phoneNumber"`
	NationalCode string `json:"nationalCode"`
}

type OTPInput struct {
	PhoneNumber string `json:"phoneNumber"`
	OTPCode     string `json:"otpCode"`
}

type LoginInput struct {
	PhoneNumber string `json:"phoneNumber"`
}

type UserData struct {
	ID          string           `json:"id"`
	FullName    string           `json:"fullName"`
	PhoneNumber string           `json:"phoneNumber"`
	UserState   entity.UserState `json:"userState"`
}
