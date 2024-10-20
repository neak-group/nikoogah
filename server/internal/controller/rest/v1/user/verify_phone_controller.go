package user

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/neak-group/nikoogah/internal/app/user/dto"
	"github.com/neak-group/nikoogah/internal/core/interface/security/session"
)

func (uc *UserController) VerifyPhone(c *gin.Context) {
	ctx := c.Request.Context()

	req := new(dto.OTPInput)
	err := c.Bind(req)
	if err != nil {
		return
	}

	otpToken, err := c.Cookie(OTPTokenKey)
	if err != nil {
		c.Error(err)
		return
	}

	req.OTPToken = otpToken

	userData, err := uc.identityService.Verify(ctx, req)

	if err != nil {
		c.Error(err)
		return
	}

	
	token, err := uc.sessionService.NewSession(ctx, userData.ID.String(), userData.FullName, session.DeviceInfo{
		UserAgent: c.Request.UserAgent(),
		IPAddress: c.Request.RemoteAddr,
	})

	if err != nil {
		return
	}

	if token.SessionID == "" {
		c.Error(fmt.Errorf("internal error: invalid token generated"))
		return
	}

	c.SetCookie("session-id", token.SessionID, int( 30 * 24 *time.Hour.Seconds()), "/", c.Request.Host, true, true)

	c.JSON(http.StatusOK, gin.H{
		"msg": "otp verified",
	})
}
