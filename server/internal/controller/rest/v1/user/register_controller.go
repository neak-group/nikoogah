package user

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/neak-group/nikoogah/internal/app/user/dto"
)

func (uc *UserController) RegisterUser(c *gin.Context) {
	ctx := c.Request.Context()

	req := new(dto.UserInput)
	err := c.Bind(req)
	if err != nil {
		return
	}

	token, err := uc.identityService.RegisterUser(ctx, req)
	if err != nil {
		c.Error(err)
		return
	}

	if token == "" {
		c.Error(fmt.Errorf("internal error: invalid token generated"))
		return
	}

	c.SetCookie(OTPTokenKey, token, int(2*time.Minute.Seconds()), "/", c.Request.Host, true, true)

	c.JSON(http.StatusOK, gin.H{
		"msg": "otp token sent",
	})
}
