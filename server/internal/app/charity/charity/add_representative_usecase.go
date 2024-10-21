package charity

import (
	"context"
	"fmt"

	"github.com/neak-group/nikoogah/internal/app/charity/charity/dto"
	"github.com/neak-group/nikoogah/internal/app/charity/charity/entity"
	"github.com/neak-group/nikoogah/internal/app/charity/charity/repository"
	"github.com/neak-group/nikoogah/internal/core/domain/base"
	"github.com/neak-group/nikoogah/internal/core/service/eventbus"
	"github.com/neak-group/nikoogah/utils/contextutils"
)

type AddRepresentativeUseCase struct {
	base.BaseUseCase
	repo repository.CharityRepository
}

type AddRepresentativeUCParams struct {
	base.UseCaseParams
	Repo repository.CharityRepository
}

func ProvideAddRepresentativeUC(params AddRepresentativeUCParams) *AddRepresentativeUseCase {
	return &AddRepresentativeUseCase{
		repo: params.Repo,
		BaseUseCase: base.BaseUseCase{
			Logger:          params.Logger,
			EventDispatcher: params.EventDispatcher,
		},
	}
}

func (uc AddRepresentativeUseCase) Execute(ctx context.Context, params *dto.AddRepresentativeParams) error {
	charity, err := uc.repo.FetchCharity(ctx, params.CharityID)
	if err != nil {
		return err
	}

	charity.Events = make([]eventbus.Event, 0)

	requesterID, err := contextutils.GetUserIDFromCtx(ctx)
	if err != nil {
		return err
	}

	manager, err := uc.repo.FindRepresentativeByUserID(ctx, requesterID)
	if err != nil {
		return err
	}

	if manager.Role != entity.Manager {
		return fmt.Errorf("unauthorized access")
	}

	repExists, err := uc.repo.FindExistingRepresentativeByUserID(ctx, params.UserID)
	if err != nil {
		return err
	}

	if repExists {
		return fmt.Errorf("representative already exists")
	}

	//TODO: check representative limits

	err = charity.AddRepresentative(params.UserID, entity.Employee)
	if err != nil {
		return err
	}

	_, err = uc.repo.SaveCharity(ctx, charity)
	if err != nil {
		return err
	}

	uc.EventDispatcher.DispatchBatch(charity.Events)

	return nil
}
