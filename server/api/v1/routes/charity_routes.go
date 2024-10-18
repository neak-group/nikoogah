package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/neak-group/nikoogah/api/auth"
	"go.uber.org/fx"
)

type CharityControllerGroup struct {
	fx.In
}

func AddCharityRoutes(parent *gin.RouterGroup, authRouter *auth.Authenticator, params CharityControllerGroup) {

}
