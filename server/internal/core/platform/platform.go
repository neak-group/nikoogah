package platform

import (
	"github.com/neak-group/nikoogah/internal/core/platform/eventbus"
	"github.com/neak-group/nikoogah/internal/core/platform/eventdispatcher"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"platform",
	fx.Provide(
		fx.Annotate(
			eventbus.ProvideEventBus,
			fx.ParamTags(`group:"event-handlers"`),
		),
		eventdispatcher.ProvideEventDispatcher,
	),
)
