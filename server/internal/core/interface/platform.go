package platform

import (
	"github.com/neak-group/nikoogah/internal/core/interface/eventbus"
	"github.com/neak-group/nikoogah/internal/core/interface/eventdispatcher"
	"github.com/neak-group/nikoogah/internal/core/interface/security/otp"
	"github.com/neak-group/nikoogah/internal/core/interface/security/session"
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
