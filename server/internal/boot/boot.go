package boot

import (
	"net/http"

	"github.com/joho/godotenv"
	"github.com/neak-group/nikoogah/api"
	"github.com/neak-group/nikoogah/internal/app"
	"github.com/neak-group/nikoogah/internal/controller"
	platform "github.com/neak-group/nikoogah/internal/services/core"
	"github.com/neak-group/nikoogah/internal/infra/httpserver"
	"github.com/neak-group/nikoogah/internal/infra/keystorefx"
	"github.com/neak-group/nikoogah/internal/infra/mongofx"
	"github.com/neak-group/nikoogah/internal/infra/telemetry"
	"github.com/neak-group/nikoogah/internal/repository"
	"github.com/neak-group/nikoogah/internal/services"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func Boot() (*fx.App, error) {
	run := fx.New(
		// Initiate Log
		telemetry.Module,
		//Load Envs
		fx.Invoke(func(logger *zap.Logger) {
			viper.SetEnvPrefix("nkg")
			env := viper.GetString("env")
			if env == "" {
				if err := godotenv.Load(".env"); err != nil {
					panic(err)
				}

				logger.Info("env initiated")
			}
			viper.AutomaticEnv()
		}),
		// Initiate Database
		mongofx.Module,
		//InitRedis
		keystorefx.Module,
		//InitStorage
		//Init platform services
		platform.Module,
		//Init Repositories
		repository.GetModule(),
		services.Module,
		app.GetModule(),
		controller.GetModule(),

		fx.Provide(
			api.ProvideHTTPRouter,
			httpserver.NewHTTPServer,
		),
		// Start HTTP Server
		fx.Invoke(func(*http.Server) {}),
	)

	return run, nil
}
