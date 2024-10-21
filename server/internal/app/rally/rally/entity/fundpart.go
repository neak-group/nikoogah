package entity

import (
	"time"

	"github.com/neak-group/nikoogah/utils/uuid"
	"github.com/shopspring/decimal"
)

type FundParticipation struct {
	Amount      decimal.Decimal    `bson:"amount"`
	VolunteerID uuid.UUID          `bson:"volunteer_id"`
	Date        time.Time          `bson:"date"`
}