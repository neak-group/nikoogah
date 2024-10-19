package dto

import (
	"github.com/neak-group/nikoogah/internal/app/user/entity"
	coreobjects "github.com/neak-group/nikoogah/internal/core/valueobjects"
	"github.com/neak-group/nikoogah/utils/uuid"
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
	OTPToken    string
}

type LoginInput struct {
	PhoneNumber string `json:"phoneNumber"`
}

type UserData struct {
	ID          uuid.UUID               `json:"id"`
	FullName    string                  `json:"fullName"`
	PhoneNumber coreobjects.PhoneNumber `json:"phoneNumber"`
	UserState   entity.UserState        `json:"userState"`
}
