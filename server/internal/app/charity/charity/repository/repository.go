package repository

import (
	"context"

	"github.com/neak-group/nikoogah/internal/app/charity/charity/entity"
	"github.com/neak-group/nikoogah/utils/uuid"
)

type CharityRepository interface {
	FindCharityTierID(ctx context.Context, name string) (uuid.UUID, error)
	FetchCharity(ctx context.Context, id uuid.UUID) (*entity.Charity, error)
	CreateCharity(ctx context.Context, charity *entity.Charity) (uuid.UUID, error)
	SaveCharity(ctx context.Context, charity *entity.Charity) (uuid.UUID, error)

	FindRepresentativeByUserID(ctx context.Context, userID uuid.UUID) (*entity.Representative, error)
	FindExistingRepresentativeByUserID(ctx context.Context, userID uuid.UUID) (bool, error)
}
