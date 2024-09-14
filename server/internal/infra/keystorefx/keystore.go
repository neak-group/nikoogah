package keystorefx

import "go.uber.org/fx"

var Module = fx.Module("keystore",
	fx.Provide(
		ProvideKeyStoreConfig,
		NewRedisClient,
		ProvideKeyStoreConn,
	),
)
