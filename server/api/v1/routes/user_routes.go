package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/neak-group/nikoogah/internal/controller/rest/v1/user"
	"go.uber.org/fx"
)

type UserControllerGroup struct {
	fx.In

	Controller user.UserController
}

func AddUserRoutes(parent *gin.RouterGroup, controllers UserControllerGroup) {
	routerGroup := parent.Group("/user")

	routerGroup.POST("/sign-up", controllers.Controller.RegisterUser)
	routerGroup.POST("/login", controllers.Controller.Login)
	routerGroup.POST("/otp", controllers.Controller.VerifyPhone)
}
