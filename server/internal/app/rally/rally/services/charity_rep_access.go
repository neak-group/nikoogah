package services

import (
	"context"

	"github.com/google/uuid"
)

type CharityAccessService interface {
	CanViewParticipation(ctx context.Context, CharityID uuid.UUID) (bool, error)
	CanAcceptParticipation(ctx context.Context, CharityID uuid.UUID) (bool, error)
}
