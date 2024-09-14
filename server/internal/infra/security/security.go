package security

import (
	"github.com/neak-group/nikoogah/internal/infra/security/otp"
	"github.com/neak-group/nikoogah/internal/infra/security/session"

	"go.uber.org/fx"
)

var Module = fx.Module("security",
	fx.Provide(
		otp.NewOTPGenerator,
		session.ProvideSessionService,
	),
)
