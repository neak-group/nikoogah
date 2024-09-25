package rally

import (
	"context"

	"github.com/neak-group/nikoogah/internal/app"
	"github.com/neak-group/nikoogah/internal/app/rally/rally/entity"
)

type GetParticipantsUseCase struct {
	app.BaseUseCase
	repo RallyRepository
}

type GetParticipantsUCParams struct {
	app.UseCaseParams
	Repo RallyRepository
}

func ProvideGetParticipantsUC(params GetParticipantsUCParams) *GetParticipantsUseCase {
	return &GetParticipantsUseCase{
		repo: params.Repo,
		BaseUseCase: app.BaseUseCase{
			Logger:          params.Logger,
			EventDispatcher: params.EventDispatcher,
		},
	}
}

func init() {
	app.RegisterUseCaseProvider(ProvideGetParticipantsUC)
}

type GetParticipantsParams struct {
}

func (uc *GetParticipantsUseCase) Execute(ctx context.Context, params GetParticipantsParams) ([]*entity.HumanParticipation, error) {
	return nil, nil
}
