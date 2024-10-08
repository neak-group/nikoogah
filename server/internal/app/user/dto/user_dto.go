package dto

import (
	"github.com/google/uuid"
	"github.com/neak-group/nikoogah/internal/app/user/entity"
	coreobjects "github.com/neak-group/nikoogah/internal/core/valueobjects"
)

type UserInput struct {
	FirstName    string
	LastName     string
	PhoneNumber  string
	NationalCode string
}

type OTPInput struct {
	PhoneNumber string
	OTPCode     string
	OTPToken    string
}

type LoginInput struct {
	PhoneNumber string
}

type UserData struct {
	ID          uuid.UUID
	FullName    string
	PhoneNumber coreobjects.PhoneNumber
	UserState   entity.UserState
}
