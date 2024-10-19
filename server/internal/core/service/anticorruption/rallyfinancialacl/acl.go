package rallyfinancialacl

import (
	"context"

	"github.com/neak-group/nikoogah/utils/uuid"
	"github.com/shopspring/decimal"
)

type FinancialServiceACL interface {
	CalculateRallyFee(ctx context.Context, FundAmount decimal.Decimal) (fee decimal.Decimal, err error)
	RequestPaymentForRallyFee(ctx context.Context, amount string, CharityID uuid.UUID)
	RequestPaymentForFundParticipation(ctx context.Context, amount string, VolunteerID uuid.UUID)
}
