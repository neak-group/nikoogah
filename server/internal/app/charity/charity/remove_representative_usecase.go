package charity

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/neak-group/nikoogah/internal/app"
	"github.com/neak-group/nikoogah/internal/app/charity/charity/entity"
	"github.com/neak-group/nikoogah/utils/contextutils"
)

type RemoveRepresentativeUseCase struct {
	app.BaseUseCase
	repo CharityRepository
}

type RemoveRepresentativeUCParams struct {
	app.UseCaseParams

	Repo CharityRepository
}

func ProvideRemoveRepresentativeUC(params RemoveRepresentativeUCParams) *RemoveRepresentativeUseCase {
	return &RemoveRepresentativeUseCase{
		repo: params.Repo,
		BaseUseCase: app.BaseUseCase{
			Logger:          params.Logger,
			EventDispatcher: params.EventDispatcher,
		},
	}
}

func init() {
	app.RegisterUseCaseProvider(ProvideRemoveRepresentativeUC)
}

type RemoveRepresentativeParams struct {
	CharityID uuid.UUID
	UserID    uuid.UUID
}

func (uc RemoveRepresentativeUseCase) Execute(ctx context.Context, params RemoveRepresentativeParams) error {
	charity, err := uc.repo.FetchCharity(params.CharityID)
	if err != nil {
		return err
	}

	requesterID, err := contextutils.GetUserIDFromCtx(ctx)
	if err != nil {
		return err
	}

	manager, err := uc.repo.FindRepresentativeByUserID(requesterID)
	if err != nil {
		return err
	}

	if manager.Role != entity.Manager {
		return fmt.Errorf("unauthorized access")
	}

	repExists, err := uc.repo.FindExistingRepresentativeByUserID(params.UserID)
	if err != nil {
		return err
	}

	if !repExists {
		return fmt.Errorf("representative does not exist")
	}

	err = charity.RemoveRepresentative(params.UserID, entity.Employee)
	if err != nil {
		return err
	}

	_, err = uc.repo.SaveCharity(charity)
	if err != nil {
		return err
	}

	for _, e := range charity.Events {
		if err := uc.EventDispatcher.Dispatch(e); err != nil {
			uc.Logger.Error(err.Error())
		}
	}

	return nil
}
