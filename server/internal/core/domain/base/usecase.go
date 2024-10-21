package base

import (
	"github.com/neak-group/nikoogah/internal/core/service/eventdispatcher"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type UseCaseParams struct {
	fx.In

	Logger          *zap.Logger
	EventDispatcher eventdispatcher.EventDispatcher
}

type BaseUseCase struct {
	Logger          *zap.Logger
	EventDispatcher eventdispatcher.EventDispatcher
}
