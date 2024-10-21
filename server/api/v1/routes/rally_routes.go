package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/neak-group/nikoogah/api/auth"
	"github.com/neak-group/nikoogah/internal/controller/rest/v1/rally/participation"
	"github.com/neak-group/nikoogah/internal/controller/rest/v1/rally/rally"
	"go.uber.org/fx"
)

type RallyControllerGroup struct {
	fx.In

	RallyController         *rally.RallyController
	ParticipationController *participation.ParticipationController
}

func AddRallyRoutes(parent *gin.RouterGroup, authRouter *auth.Authenticator, params RallyControllerGroup) {
	rallyGroup := parent.Group("/rally")

	rallyGroup.POST("/", params.RallyController.CreateRally)
	rallyGroup.GET("/", params.RallyController.FetchRallies)
	rallyGroup.GET("/:rally-id", params.RallyController.FetchRally)

	participationGroup := rallyGroup.Group("/participation")

	participationGroup.PUT("/fund", params.ParticipationController.RegisterFundParticipation)
	participationGroup.PUT("/aid", params.ParticipationController.RegisterHumanParticipation)
	participationGroup.GET("/aid", params.ParticipationController.FetchParticipants)



}
