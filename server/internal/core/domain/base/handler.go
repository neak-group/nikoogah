package base

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type HandlerParams struct {
	fx.In
	
	Logger *zap.Logger
}

type BaseHandler struct {
	Logger *zap.Logger
}
