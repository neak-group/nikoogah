package api

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/neak-group/nikoogah/api/auth"
	v1 "github.com/neak-group/nikoogah/api/v1"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type HTTPRouterParams struct {
	fx.In

	v1.V1RoutesParams
	Logger *zap.Logger
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
	e.Use(HandleError())
	apiGroup := e.Group("/api")

	authenticator := auth.NewAuthenticator(params.SessionService)
	apiGroup.Use(authenticator.Authenticate(params.Logger))

	v1.AddV1Routes(apiGroup, authenticator, params.V1RoutesParams)

	return e
}

func HandleError() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 {
			return
		}
		// ctx := c.Request.Context()

		for _, ginerr := range c.Errors {
			c.JSON(400, ginerr.Err.Error())
			return
		}

	}
}
