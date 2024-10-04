package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/neak-group/nikoogah/internal/app/rally/volunteer/entity"
)

type VolunteerRepository interface {
	FetchVolunteersByBatchID(ctx context.Context, ids uuid.UUIDs) ([]*entity.Volunteer, error)
}
