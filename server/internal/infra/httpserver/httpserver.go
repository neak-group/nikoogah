package httpserver

import (
	"context"
	"net"
	"net/http"

	"github.com/spf13/viper"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewHTTPServer(lc fx.Lifecycle, mux http.Handler, log *zap.Logger) *http.Server {
	listenAddress := viper.GetString("http_listen_address")
	srv := &http.Server{Addr: listenAddress, Handler: mux}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			lsn, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				log.Error(err.Error())
				return err
			}

			log.Info("Starting HTTP server on", zap.String("addr", srv.Addr))

			go srv.Serve(lsn)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})

	return srv
}
