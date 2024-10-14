package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/neak-group/nikoogah/internal/controller/rest/v1/user"
)

type UserControllerGroup struct {
	controller user.UserController
}

func AddUserRoutes(parent *gin.RouterGroup, controllers UserControllerGroup) {
	routerGroup := parent.Group("/user")

	routerGroup.POST("/sign-up", controllers.controller.RegisterUser)
	routerGroup.POST("/login", controllers.controller.Login)
	routerGroup.POST("/otp", controllers.controller.VerifyPhone)
}
