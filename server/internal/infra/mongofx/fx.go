package mongofx

import "go.uber.org/fx"

var Module = fx.Module("db",
	fx.Provide(
		ProvideMongoDBConfig,
		NewMongoConn,
		NewMongoDBConn,
	),
)
