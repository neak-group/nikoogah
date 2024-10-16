package rest

import (
	v1 "github.com/neak-group/nikoogah/internal/controller/rest/v1"
	"go.uber.org/fx"
)

func GetModule() fx.Option {
	return fx.Module("application",
		fx.Provide(v1.ProvideV1RestControllers()...),
	)
}
