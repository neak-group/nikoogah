package repository

import (
	"context"

	"github.com/neak-group/nikoogah/internal/app/rally/volunteer/entity"
	"github.com/neak-group/nikoogah/utils/uuid"
)

type VolunteerRepository interface {
	UpdateVolunteer(ctx context.Context, volunteer *entity.Volunteer) error
	FetchVolunteer(ctx context.Context, id uuid.UUID) (*entity.Volunteer, error)
	FetchVolunteersByBatchID(ctx context.Context, ids []uuid.UUID) ([]*entity.Volunteer, error)
}
