package mongofx

import "go.uber.org/fx"

var Module = fx.Module("mongodb",
	fx.Provide(
		ProvideMongoDBConfig,
		NewMongoConn,
		NewMongoDBConn,
	),
)
