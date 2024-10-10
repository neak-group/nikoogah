package boot

import (
	"net/http"

	"github.com/joho/godotenv"
	"github.com/neak-group/nikoogah/internal/app"
	"github.com/neak-group/nikoogah/internal/controller"
	platform "github.com/neak-group/nikoogah/internal/core/interface"
	"github.com/neak-group/nikoogah/internal/infra/dbfx"
	"github.com/neak-group/nikoogah/internal/infra/httpserver"
	"github.com/neak-group/nikoogah/internal/infra/keystorefx"
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
			viper.SetEnvPrefix("nkg")
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
		//Init platform services
		platform.Module,
		//Init Repositories
		repository.GetModule(),
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
