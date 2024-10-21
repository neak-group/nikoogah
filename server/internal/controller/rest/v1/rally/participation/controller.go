package participation

import (
	"github.com/neak-group/nikoogah/internal/app/rally/rally"
	v1 "github.com/neak-group/nikoogah/internal/controller/rest/v1"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type ParticipationControllerParams struct {
	fx.In

	GetParticipantsUseCase       *rally.GetParticipantsUseCase
	NewHumanParticipationUseCase *rally.NewHumanParticipationUseCase
	NewFundParticipationUseCase  *rally.NewFundParticipationUseCase
	Logger                       *zap.Logger
}

type ParticipationController struct {
	getParticipantsUseCase       *rally.GetParticipantsUseCase
	newHumanParticipationUseCase *rally.NewHumanParticipationUseCase
	newFundParticipationUseCase  *rally.NewFundParticipationUseCase
	logger                       *zap.Logger
}

func NewParticipationController(params ParticipationControllerParams) *ParticipationController {
	return &ParticipationController{
		getParticipantsUseCase:       params.GetParticipantsUseCase,
		newHumanParticipationUseCase: params.NewHumanParticipationUseCase,
		newFundParticipationUseCase:  params.NewFundParticipationUseCase,
		logger:                       params.Logger,
	}
}

func init() {
	v1.RegisterControllerProvider(NewParticipationController)
}
