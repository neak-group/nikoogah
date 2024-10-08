package user

import (
	"github.com/neak-group/nikoogah/internal/app"
	"github.com/neak-group/nikoogah/internal/app/user/repository"
	"github.com/neak-group/nikoogah/internal/app/user/services"
	"github.com/neak-group/nikoogah/internal/core/service/eventdispatcher"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type IdentityService struct {
	userRepo        repository.UserRepository
	logger          *zap.Logger
	eventDispatcher eventdispatcher.EventDispatcher
	otpService      services.OTPService
}

type IdentityServiceParams struct {
	fx.In

	UserRepo        repository.UserRepository
	Logger          *zap.Logger
	EventDispatcher eventdispatcher.EventDispatcher
	OTPService      services.OTPService
	SessionService services.SessionService
}

func ProvideIdentityService(params IdentityServiceParams) *IdentityService {
	return &IdentityService{
		userRepo:        params.UserRepo,
		logger:          params.Logger,
		eventDispatcher: params.EventDispatcher,
		otpService:      params.OTPService,
	}
}

func init() {
	app.RegisterUseCaseProvider(ProvideIdentityService)
}
