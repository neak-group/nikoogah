package platform

import (
	"github.com/neak-group/nikoogah/internal/core/interface/eventbus"
	"github.com/neak-group/nikoogah/internal/core/interface/eventdispatcher"
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
