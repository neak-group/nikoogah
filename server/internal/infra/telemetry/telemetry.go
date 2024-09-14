package telemetry

import (
	"github.com/neak-group/nikoogah/internal/infra/telemetry/logging"

	"go.uber.org/fx"
)

var Module = fx.Module("telemetry",
	logging.Module,
)
