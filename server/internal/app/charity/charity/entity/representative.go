package entity

import (
	"time"

	"github.com/neak-group/nikoogah/utils/uuid"
)

type Representative struct {
	UserID   uuid.UUID          `bson:"user_id"`
	Role     RepresentativeRole `bson:"role"`
	JoinedAt time.Time          `bson:"joined_at"`
}

type RepresentativeRole string

const (
	Manager  RepresentativeRole = "manager"
	Employee RepresentativeRole = "employee"
)
