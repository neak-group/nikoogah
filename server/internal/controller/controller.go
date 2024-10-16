package controller

import (
	"github.com/neak-group/nikoogah/internal/controller/rest"
	"go.uber.org/fx"
)

func GetModule() fx.Option {
	return fx.Module("controller", rest.GetModule())
}
