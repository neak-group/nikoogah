package entity

import (
	"time"

	"github.com/google/uuid"
)

type Representative struct {
	UserID   uuid.UUID
	Role     RepresentativeRole
	JoinedAt time.Time
}

type RepresentativeRole string

const (
	Manager  RepresentativeRole = "manager"
	Employee RepresentativeRole = "employee"
)
