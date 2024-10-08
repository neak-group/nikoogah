package user

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/neak-group/nikoogah/internal/app/user"
	"github.com/neak-group/nikoogah/internal/app/user/dto"
	"github.com/neak-group/nikoogah/internal/controller/rest/v1/user/model"
	"github.com/neak-group/nikoogah/internal/infra/security/session"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

const OTPTokenKey = "otp-token"

type UserHandler struct {
	identityService *user.IdentityService
	sessionService  *session.SessionService

	logger *zap.Logger
}

type UserHandlerParams struct {
	fx.In

	IdentityService *user.IdentityService
	SessionService  *session.SessionService

	Logger *zap.Logger
}

func NewUserController(params UserHandlerParams) UserHandler {
	return UserHandler{
		identityService: params.IdentityService,
		logger:          params.Logger,
	}
}

func (uc *UserHandler) RegisterUser(c *gin.Context) {
	ctx := c.Request.Context()

	req := new(model.UserInput)
	err := c.Bind(req)
	if err != nil {
		return
	}

	token, err := uc.identityService.RegisterUser(ctx, dto.UserInput{
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		PhoneNumber:  req.PhoneNumber,
		NationalCode: req.NationalCode,
	})
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

func (uc *UserHandler) VerifyRegistration(c *gin.Context) {
	ctx := c.Request.Context()

	req := new(model.OTPInput)
	err := c.Bind(req)
	if err != nil {
		return
	}

	otpToken, err := c.Cookie(OTPTokenKey)
	if err != nil {
		return
	}

	sessionData, err := uc.identityService.VerifyRegistration(ctx, dto.OTPInput{
		PhoneNumber: req.PhoneNumber,
		OTPCode:     req.OTPCode,
		OTPToken:    otpToken,
	})

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

	if token == "" {
		c.Error(fmt.Errorf("internal error: invalid token generated"))
		return
	}

	c.SetCookie("otp-token", token, int(2*time.Minute.Seconds()), "/", c.Request.Host, true, true)

	c.JSON(http.StatusOK, gin.H{
		"msg": "otp token sent",
	})
}

func (uc *UserHandler) Login(c *gin.Context) {

}

func (uc *UserHandler) VerifyLogin(c *gin.Context) {

}
