package repository

import (
	"context"

	"github.com/neak-group/nikoogah/internal/app/rally/volunteer/entity"
	"github.com/neak-group/nikoogah/utils/uuid"
)

type VolunteerRepository interface {
	FetchVolunteersByBatchID(ctx context.Context, ids uuid.UUIDs) ([]*entity.Volunteer, error)
}
