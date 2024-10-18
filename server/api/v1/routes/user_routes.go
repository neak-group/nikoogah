package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/neak-group/nikoogah/api/auth"
	"github.com/neak-group/nikoogah/internal/controller/rest/v1/user"
	"go.uber.org/fx"
)

type UserControllerGroup struct {
	fx.In

	Controller user.UserController
}

func AddUserRoutes(parent *gin.RouterGroup, authRouter *auth.Authenticator, controllers UserControllerGroup) {
	routerGroup := parent.Group("/user")

	routerGroup.POST("/sign-up", controllers.Controller.RegisterUser)
	authRouter.AddAnonymousRoute(http.MethodPost, routerGroup.BasePath()+"/sign-up")

	routerGroup.POST("/login", controllers.Controller.Login)
	authRouter.AddAnonymousRoute(http.MethodPost, routerGroup.BasePath()+"/login")

	routerGroup.POST("/otp", controllers.Controller.VerifyPhone)
	authRouter.AddAnonymousRoute(http.MethodPost, routerGroup.BasePath()+"/otp")

}
