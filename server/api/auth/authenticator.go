package auth

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/neak-group/nikoogah/internal/core/interface/security/session"
	"github.com/neak-group/nikoogah/utils/contextutils"
	"go.uber.org/zap"
)

type Authenticator struct {
	sessionService  session.SessionService
	anonymousRoutes map[string][]string
}

func (r Authenticator) Authenticate(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, r := range r.anonymousRoutes[c.FullPath()] {
			if r == c.Request.Method {
				logger.Info("route bypassed", zap.String("route", c.FullPath()))
				c.Next()
				return
			}
		}

		ctx := c.Request.Context()
		if ctx == nil {
			c.Error(fmt.Errorf("context does not exist"))
			return
		}

		authCookie, err := c.Request.Cookie("session-id")
		if err != nil {
			c.Error(err)
			return
		}

		sessionID := authCookie.Value

		session, err := r.sessionService.ValidateSession(ctx, sessionID)
		if err != nil {
			c.Error(err)
			return
		}

		ctx = contextutils.SetUserIDCtx(ctx, session.UserID)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
