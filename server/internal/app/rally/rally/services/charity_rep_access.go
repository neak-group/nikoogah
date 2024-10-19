package services

import (
	"context"

	"github.com/neak-group/nikoogah/utils/uuid"
)

type CharityAccessService interface {
	CanViewParticipation(ctx context.Context, CharityID uuid.UUID) (bool, error)
	CanAcceptParticipation(ctx context.Context, CharityID uuid.UUID) (bool, error)
}
