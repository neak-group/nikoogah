package routes

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type RallyControllerGroup struct {
	fx.In
}

func AddRallyRoutes(parent *gin.RouterGroup, params RallyControllerGroup) {

}
