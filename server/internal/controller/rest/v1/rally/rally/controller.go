package rally

import (
	"github.com/neak-group/nikoogah/internal/app/rally/rally"
	v1 "github.com/neak-group/nikoogah/internal/controller/rest/v1"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type RallyControllerParams struct {
	fx.In

	FetchRallyUseCase          *rally.FetchRallyUseCase
	FetchRalliesUseCase        *rally.FetchRalliesUseCase
	FetchCharityRalliesUseCase *rally.FetchCharityRalliesUseCase
	NewRallyUseCase            *rally.NewRallyUseCase
	Logger                     *zap.Logger
}

type RallyController struct {
	fetchRallyUseCase          *rally.FetchRallyUseCase
	fetchRalliesUseCase        *rally.FetchRalliesUseCase
	fetchCharityRalliesUseCase *rally.FetchCharityRalliesUseCase
	newRallyUseCase            *rally.NewRallyUseCase
	logger                     *zap.Logger
}

func NewRallyController(params RallyControllerParams) *RallyController {
	return &RallyController{
		fetchRallyUseCase:          params.FetchRallyUseCase,
		fetchRalliesUseCase:        params.FetchRalliesUseCase,
		fetchCharityRalliesUseCase: params.FetchCharityRalliesUseCase,
		newRallyUseCase:            params.NewRallyUseCase,
		logger:                     params.Logger,
	}
}

func init() {
	v1.RegisterControllerProvider(NewRallyController)
}
