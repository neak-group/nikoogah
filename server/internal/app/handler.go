package app

import (
	"github.com/neak-group/nikoogah/internal/core/service/eventbus"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func RegisterHandlerProvider(provider interface{}) {
	if provider == nil {
		return
	}

	if eventHandlerProviders == nil {
		eventHandlerProviders = append(eventHandlerProviders, fx.Annotate(
			provider,
			fx.As(new(eventbus.EventHandler)),
			fx.ResultTags(`group:"event-handlers"`),
		))
	}
}

type HandlerParams struct {
	Logger *zap.Logger
}

type BaseHandler struct {
	Logger *zap.Logger
}
