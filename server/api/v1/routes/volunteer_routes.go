package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/neak-group/nikoogah/api/auth"
	"github.com/neak-group/nikoogah/internal/controller/rest/v1/rally/volunteer"
	"go.uber.org/fx"
)

type VolunteerControllerGroup struct {
	fx.In

	VolunteerController *volunteer.VolunteerController
}

func AddVolunteerRoutes(parent *gin.RouterGroup, authRouter *auth.Authenticator, params VolunteerControllerGroup) {

	volunteerGroup := parent.Group("/volunteer")

	volunteerGroup.GET("/profile", params.VolunteerController.FetchProfile)
}
