package api

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	v1 "github.com/neak-group/nikoogah/api/v1"
	"go.uber.org/fx"
)

type HTTPRouterParams struct {
	fx.In

	v1.V1RoutesParams
}

func ProvideHTTPRouter(params HTTPRouterParams) http.Handler {
	e := gin.Default()

	corsHandler := cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	})

	e.Use(corsHandler)

	apiGroup := e.Group("/api")

	v1.AddV1Routes(apiGroup, params.V1RoutesParams)

	return e
}
