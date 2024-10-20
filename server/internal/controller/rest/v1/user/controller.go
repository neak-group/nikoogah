package user

import (
	"github.com/neak-group/nikoogah/internal/app/user"
	v1 "github.com/neak-group/nikoogah/internal/controller/rest/v1"
	"github.com/neak-group/nikoogah/internal/core/interface/security/session"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

const OTPTokenKey = "otp-token"

type UserController struct {
	identityService *user.IdentityService
	sessionService  *session.SessionService

	logger *zap.Logger
}

type UserControllerParams struct {
	fx.In

	IdentityService *user.IdentityService
	SessionService  *session.SessionService

	Logger *zap.Logger
}

func NewUserController(params UserControllerParams) *UserController {
	return &UserController{
		identityService: params.IdentityService,
		sessionService:  params.SessionService,
		logger:          params.Logger,
	}
}

func init() {
	v1.RegisterControllerProvider(NewUserController)
}
