package representative

import (
	"github.com/neak-group/nikoogah/internal/app/charity/charity"
	v1 "github.com/neak-group/nikoogah/internal/controller/rest/v1"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type RepresentativeControllerParams struct {
	fx.In

	AddRepUseCase    *charity.AddRepresentativeUseCase
	RemoveRepUseCase *charity.RemoveRepresentativeUseCase

	Logger *zap.Logger
}

type RepresentativeController struct {
	addRepUseCase    *charity.AddRepresentativeUseCase
	removeRepUseCase *charity.RemoveRepresentativeUseCase
	logger           *zap.Logger
}

func NewRepresentativeController(params RepresentativeControllerParams) *RepresentativeController {
	return &RepresentativeController{
		addRepUseCase:    params.AddRepUseCase,
		removeRepUseCase: params.RemoveRepUseCase,
		logger:           params.Logger,
	}
}

func init() {
	v1.RegisterControllerProvider(NewRepresentativeController)
}
