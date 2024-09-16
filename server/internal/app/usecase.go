package app

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type UseCaseParams struct {
	fx.In

	Logger *zap.Logger
}


type BaseUseCase struct{
	Logger *zap.Logger
}