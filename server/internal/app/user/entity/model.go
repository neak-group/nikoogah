package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/neak-group/nikoogah/internal/core/service/eventbus"
	coreobjects "github.com/neak-group/nikoogah/internal/core/valueobjects"
)

type UserState int

const (
	UserPending UserState = iota
	UserActive
	UserSuspended
)

type User struct {
	ID uuid.UUID

	FirstName       string
	LastName        string
	PhoneNumber     coreobjects.PhoneNumber
	PhoneVerifiedAt *time.Time
	NationalCode    string
	AvatarPath      string
	ResumePath      string
	UserState       UserState

	CreatedAt time.Time
	UpdatedAt time.Time

	Events []eventbus.Event
}

func NewUser(firstName, lastName, phoneNumber, nationalCode string) (*User, error) {
	return &User{
		ID:        uuid.New(),
		FirstName: firstName,
		LastName:  lastName,
		PhoneNumber: coreobjects.PhoneNumber{
			PhoneNumber: phoneNumber,
			Region:      "IR",
		},
		NationalCode: nationalCode,
		UserState:    UserPending,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}, nil
}
