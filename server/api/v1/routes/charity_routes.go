package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/neak-group/nikoogah/api/auth"
	"github.com/neak-group/nikoogah/internal/controller/rest/v1/charity/charity"
	"github.com/neak-group/nikoogah/internal/controller/rest/v1/charity/representative"
	"go.uber.org/fx"
)

type CharityControllerGroup struct {
	fx.In

	CharityController        *charity.CharityController
	RepresentativeController *representative.RepresentativeController
}

func AddCharityRoutes(parent *gin.RouterGroup, authRouter *auth.Authenticator, params CharityControllerGroup) {
	charityGroup := parent.Group("/charity")

	charityGroup.GET("/:charity-id", params.CharityController.FetchCharity)
	charityGroup.POST("/", params.CharityController.RegisterCharity)

	repGroup := charityGroup.Group("/:charity-id/rep")
	repGroup.PUT("/", params.RepresentativeController.AddRepresentative)
	repGroup.DELETE("/rep-id", params.RepresentativeController.RemoveRepresentative)
}
