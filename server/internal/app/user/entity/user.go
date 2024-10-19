package entity

import (
	"time"

	"github.com/neak-group/nikoogah/internal/core/service/eventbus"
	coreobjects "github.com/neak-group/nikoogah/internal/core/valueobjects"
	"github.com/neak-group/nikoogah/utils/uuid"
)

type UserState int

const (
	UserPending UserState = iota
	UserActive
	UserSuspended
)

type User struct {
	ID uuid.UUID `bson:"id"`

	FirstName       string                  `bson:"first_name"`
	LastName        string                  `bson:"last_name"`
	PhoneNumber     coreobjects.PhoneNumber `bson:"phone_number"`
	PhoneVerifiedAt *time.Time              `bson:"phone_verified_at"`
	NationalCode    string                  `bson:"national_code"`
	AvatarPath      string                  `bson:"avatar_path"`
	ResumePath      string                  `bson:"resume_path"`
	UserState       UserState               `bson:"user_state"`

	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`

	Events []eventbus.Event `bson:"-"`
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
