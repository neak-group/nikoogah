package repository

import (
	"context"

	"github.com/neak-group/nikoogah/internal/app/volunteer/volunteer/entity"
	"github.com/neak-group/nikoogah/utils/uuid"
)

type VolunteerRepository interface {
	UpdateRepository(ctx context.Context, volunteer *entity.Volunteer) error
	FetchVolunteer(ctx context.Context, id uuid.UUID) (*entity.Volunteer, error)
}
