package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/neak-group/nikoogah/api/auth"
	"go.uber.org/fx"
)

type RallyControllerGroup struct {
	fx.In
}

func AddRallyRoutes(parent *gin.RouterGroup, authRouter *auth.Authenticator, params RallyControllerGroup) {

}
