package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/neak-group/nikoogah/api/v1/routes"
)

type V1RoutesParams struct {
	routes.UserControllerGroup
	routes.CharityControllerGroup
	routes.RallyControllerGroup
}

func AddV1Routes(baseRouter *gin.RouterGroup, params V1RoutesParams) {

	apiv1 := baseRouter.Group("/v1")

	routes.AddUserRoutes(apiv1, params.UserControllerGroup)
	routes.AddRallyRoutes(apiv1, params.RallyControllerGroup)
	routes.AddCharityRoutes(apiv1, params.CharityControllerGroup)
}
