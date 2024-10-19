package charity

import (
	"context"

	"github.com/neak-group/nikoogah/internal/app"
	"github.com/neak-group/nikoogah/internal/app/charity/charity/dto"
	"github.com/neak-group/nikoogah/internal/app/charity/charity/entity"
	"github.com/neak-group/nikoogah/internal/app/charity/charity/repository"
	"github.com/neak-group/nikoogah/internal/core/service/eventbus"
)

type CheckRepresentativeAccessUseCase struct {
	app.BaseUseCase
	repo repository.CharityRepository
}

type CheckRepresentativeAccessUCParams struct {
	app.UseCaseParams

	Repo repository.CharityRepository
}

func ProvideCheckRepresentativeAccessUC(params CheckRepresentativeAccessUCParams) *CheckRepresentativeAccessUseCase {
	return &CheckRepresentativeAccessUseCase{
		repo: params.Repo,
		BaseUseCase: app.BaseUseCase{
			Logger:          params.Logger,
			EventDispatcher: params.EventDispatcher,
		},
	}
}

func init() {
	app.RegisterUseCaseProvider(ProvideCheckRepresentativeAccessUC)
}

func (uc CheckRepresentativeAccessUseCase) Execute(ctx context.Context, params dto.CheckRepresentativeAccessParams) (bool, error) {
	charity, err := uc.repo.FetchCharity(ctx, params.CharityID)
	if err != nil {
		return false, err
	}
	charity.Events = make([]eventbus.Event, 0)

	rep, err := uc.repo.FindRepresentativeByUserID(ctx, params.UserID)
	if err != nil {
		return false, err
	}

	AK, err := entity.MapAccessKey(params.AccessKey)
	if err != nil {
		return false, err
	}

	accesskeys := entity.GetRoleAccess(rep.Role)

	for _, key := range accesskeys {
		if key == AK {
			return true, nil
		}
	}

	return false, nil
}
