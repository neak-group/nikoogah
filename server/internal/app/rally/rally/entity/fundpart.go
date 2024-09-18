package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type FundParticipation struct {
	Amount      decimal.Decimal
	VolunteerID uuid.UUID
	Date        time.Time
}