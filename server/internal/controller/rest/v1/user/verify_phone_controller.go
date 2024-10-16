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
		return
	}

	req.OTPToken = otpToken

	sessionData, err := uc.identityService.Verify(ctx, req)

	if err != nil {
		c.Error(err)
		return
	}

	token, err := uc.sessionService.NewSession(ctx, sessionData.ID.String(), sessionData.FullName, session.DeviceInfo{
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

	c.SetCookie("otp-token", token.SessionID, int(2*time.Minute.Seconds()), "/", c.Request.Host, true, true)

	c.JSON(http.StatusOK, gin.H{
		"msg": "otp token sent",
	})
}
