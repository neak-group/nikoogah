package charity

import (
	"github.com/neak-group/nikoogah/internal/app/charity/charity"
	v1 "github.com/neak-group/nikoogah/internal/controller/rest/v1"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type CharityControllerParams struct {
	fx.In

	RegisterUseCase      *charity.RegisterCharityUseCase
	FetchCharityUseCase  *charity.FetchCharityUseCase
	ModifyCharityUseCase *charity.ModifyCharityUseCase

	Logger *zap.Logger
}

type CharityController struct {
	registerUseCase      *charity.RegisterCharityUseCase
	fetchCharityUseCase  *charity.FetchCharityUseCase
	modifyCharityUseCase *charity.ModifyCharityUseCase

	logger *zap.Logger
}

func NewCharityController(params CharityControllerParams) *CharityController {
	return &CharityController{
		registerUseCase:      params.RegisterUseCase,
		fetchCharityUseCase:  params.FetchCharityUseCase,
		modifyCharityUseCase: params.ModifyCharityUseCase,
		logger:               params.Logger,
	}
}

func init() {
	v1.RegisterControllerProvider(NewCharityController)
}
