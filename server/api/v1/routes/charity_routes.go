package routes

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type CharityControllerGroup struct {
	fx.In
}

func AddCharityRoutes(parent *gin.RouterGroup, params CharityControllerGroup) {

}
