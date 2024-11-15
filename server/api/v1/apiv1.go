package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/neak-group/nikoogah/api/auth"
	"github.com/neak-group/nikoogah/api/v1/routes"
	"github.com/neak-group/nikoogah/internal/services/core/security/session"
	"go.uber.org/fx"
)

type V1RoutesParams struct {
	fx.In

	SessionService *session.SessionService

	routes.UserControllerGroup
	routes.CharityControllerGroup
	routes.RallyControllerGroup
	routes.VolunteerControllerGroup
}

func AddV1Routes(baseRouter *gin.RouterGroup, authenticator *auth.Authenticator, params V1RoutesParams) {

	apiv1 := baseRouter.Group("/v1")

	routes.AddUserRoutes(apiv1, authenticator, params.UserControllerGroup)
	routes.AddRallyRoutes(apiv1, authenticator, params.RallyControllerGroup)
	routes.AddCharityRoutes(apiv1, authenticator, params.CharityControllerGroup)
	routes.AddVolunteerRoutes(apiv1, authenticator, params.VolunteerControllerGroup)
}
