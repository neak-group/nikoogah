package platform

import (
	"github.com/neak-group/nikoogah/internal/core/platform/eventbus"
	"github.com/neak-group/nikoogah/internal/core/platform/eventdispatcher"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"platform",
	fx.Provide(
		eventbus.ProvideEventBus,
		eventdispatcher.ProvideEventDispatcher,
	),
)
