package repository

import (
	"context"

	"github.com/neak-group/nikoogah/internal/app/rally/charity/entity"
	"github.com/neak-group/nikoogah/utils/uuid"
)

type CharityRepository interface {
	FetchCharity(ctx context.Context, charityID uuid.UUID) (*entity.Charity, error)
	SaveCharity(ctx context.Context, charity *entity.Charity) error
}
