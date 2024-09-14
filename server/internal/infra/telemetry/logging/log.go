package logging

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

var Module = fx.Module(
	"logging",
	fx.Provide(
		zap.NewExample,
	),
	fx.Invoke(func(*zap.Logger) {}),
	fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
		return &fxevent.ZapLogger{Logger: log}
	}),
)
