package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/neak-group/nikoogah/api/auth"
	"go.uber.org/fx"
)

type FinancialControllerGroup struct {
	fx.In
}

func AddFinancialRoutes(parent *gin.RouterGroup, authRouter *auth.Authenticator, controller FinancialControllerGroup) {

}
