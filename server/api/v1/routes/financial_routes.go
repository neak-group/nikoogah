package routes

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type FinancialControllerGroup struct {
	fx.In
}

func AddFinancialRoutes(parent *gin.RouterGroup, controller FinancialControllerGroup) {

}
