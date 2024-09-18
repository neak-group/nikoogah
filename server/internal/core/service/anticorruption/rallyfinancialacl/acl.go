package rallyfinancialacl

import (
	"context"

	"github.com/google/uuid"
)

type FinancialServiceACL interface {
	RequestPaymentForRallyFee(ctx context.Context, amount string, CharityID uuid.UUID)
	RequestPaymentForFundParticipation(ctx context.Context, amount string, VolunteerID uuid.UUID)
}
