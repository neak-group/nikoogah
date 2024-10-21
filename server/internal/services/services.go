package services

import (
	"github.com/neak-group/nikoogah/internal/services/charityaccess"
	"go.uber.org/fx"
)

var Module = fx.Module("services",
	fx.Option(
		fx.Provide(charityaccess.ProvideCharityAccessService),
	),
)
