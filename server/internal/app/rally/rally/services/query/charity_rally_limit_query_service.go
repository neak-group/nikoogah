package query

import (
	"context"
	"errors"

	"github.com/neak-group/nikoogah/internal/app/rally/charity/repository"
	rallyRepo "github.com/neak-group/nikoogah/internal/app/rally/rally/repository"
	"github.com/neak-group/nikoogah/utils/uuid"
	"go.uber.org/fx"
)

type CharityRallyQueryService interface {
	CheckCharityRallyLimit(ctx context.Context, charityID uuid.UUID) (bool, error)
}

type CharityRallyQueryServiceImpl struct {
	charityRepo repository.CharityRepository
	rallyRepo   rallyRepo.RallyRepository
}

type CharityRallyQueryServiceParams struct {
	fx.In

	CharityRepo repository.CharityRepository
	RallyRepo   rallyRepo.RallyRepository
}

func NewCharityRallyQueryService(params CharityRallyQueryServiceParams) CharityRallyQueryService {
	return &CharityRallyQueryServiceImpl{
		charityRepo: params.CharityRepo,
		rallyRepo:   params.RallyRepo,
	}
}

func (qs *CharityRallyQueryServiceImpl) CheckCharityRallyLimit(ctx context.Context, charityID uuid.UUID) (bool, error) {
	// Fetch charity details to get the rally limit
	charity, err := qs.charityRepo.FetchCharity(ctx, charityID)
	if err != nil {
		return false, err
	}

	if charity == nil {
		return false, errors.New("charity not found")
	}

	// Fetch the current number of rallies for the charity
	rallyCount, err := qs.rallyRepo.FetchCharityRallyCount(ctx, charityID)
	if err != nil {
		return false, err
	}

	// Check if the rally count exceeds the limit
	if rallyCount >= charity.MaxRallyLimit {
		return true, nil // Limit exceeded
	}

	return false, nil // Limit not exceeded
}
