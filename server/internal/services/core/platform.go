package platform

import (
	"github.com/neak-group/nikoogah/internal/services/core/eventbus"
	"github.com/neak-group/nikoogah/internal/services/core/eventdispatcher"
	"github.com/neak-group/nikoogah/internal/services/core/security/otp"
	"github.com/neak-group/nikoogah/internal/services/core/security/session"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"platform",
	fx.Provide(
		eventbus.ProvideEventBus,
		eventdispatcher.ProvideEventDispatcher,
	),
	fx.Provide(
		otp.NewOTPGenerator,
		session.ProvideSessionService,
	),
)
