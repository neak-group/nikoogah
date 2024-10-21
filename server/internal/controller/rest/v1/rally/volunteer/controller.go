package volunteer

import (
	"github.com/neak-group/nikoogah/internal/app/rally/volunteer"
	v1 "github.com/neak-group/nikoogah/internal/controller/rest/v1"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type VolunteerControllerParams struct {
	fx.In

	FetchProfileUseCase *volunteer.FetchProfileUseCase
	Logger              *zap.Logger
}

type VolunteerController struct {
	fetchProfileUseCase *volunteer.FetchProfileUseCase
	logger              *zap.Logger
}

func NewVolunteerController(params VolunteerControllerParams) *VolunteerController {
	return &VolunteerController{
		fetchProfileUseCase: params.FetchProfileUseCase,
		logger:              params.Logger,
	}
}

func init() {
	v1.RegisterControllerProvider(NewVolunteerController)
}
