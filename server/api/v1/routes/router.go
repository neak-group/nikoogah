package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HTTPRouter struct {
	engine *gin.Engine
}

type HTTPRouterParams struct {
	UserRoutes

	CharityRoutes
}

type HTTPRoute struct {
	Path    string
	Handler gin.HandlerFunc
}

func NewHTTPRouter() http.Handler {
	
}
