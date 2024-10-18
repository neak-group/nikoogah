package app

import (
	"github.com/neak-group/nikoogah/internal/core/service/eventdispatcher"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func RegisterUseCaseProvider(provider interface{}) {
	if provider == nil {
		return
	}

	if useCaseProviders == nil {
		useCaseProviders = make([]interface{}, 0)
	}
	useCaseProviders = append(useCaseProviders, provider)
}

type UseCaseParams struct {
	fx.In

	Logger          *zap.Logger
	EventDispatcher eventdispatcher.EventDispatcher
}

type BaseUseCase struct {
	Logger          *zap.Logger
	EventDispatcher eventdispatcher.EventDispatcher
}
