package boot

import (
	"net/http"

	"github.com/joho/godotenv"
	"github.com/neak-group/nikoogah/internal/app"
	"github.com/neak-group/nikoogah/internal/controller"
	"github.com/neak-group/nikoogah/internal/core/platform/eventbus"
	"github.com/neak-group/nikoogah/internal/infra/dbfx"
	"github.com/neak-group/nikoogah/internal/infra/httpserver"
	"github.com/neak-group/nikoogah/internal/infra/keystorefx"
	"github.com/neak-group/nikoogah/internal/infra/security"
	"github.com/neak-group/nikoogah/internal/infra/telemetry"
	"github.com/neak-group/nikoogah/internal/repository"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

func Boot() (*fx.App, error) {
	run := fx.New(
		// Initiate Log
		telemetry.Module,
		//Load Envs
		fx.Invoke(func() {
			viper.SetEnvPrefix("ghdm")
			viper.AutomaticEnv()
			env := viper.GetString("env")
			if env == "" {
				if err := godotenv.Load(".env"); err != nil {
					panic(err)
				}
			}
		}),
		// Initiate Database
		dbfx.Module,
		//InitRedis
		keystorefx.Module,
		//InitStorage
		security.Module,
		//Init Repositories
		repository.GetModule(),
		eventbus.Module,
		app.GetModule(),
		controller.GetModule(),
		fx.Provide(
			httpserver.NewHTTPServer,
		),
		// Start HTTP Server
		fx.Invoke(func(*http.Server) {}),
	)

	return run, nil
}
