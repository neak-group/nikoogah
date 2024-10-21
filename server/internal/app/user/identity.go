package user

import (
	"github.com/neak-group/nikoogah/internal/app/user/repository"
	"github.com/neak-group/nikoogah/internal/core/service/eventdispatcher"
	"github.com/neak-group/nikoogah/internal/core/service/security/otp"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type IdentityService struct {
	userRepo        repository.UserRepository
	logger          *zap.Logger
	eventDispatcher eventdispatcher.EventDispatcher
	otpService      otp.OTPService
}

func GetUseCaseProviders() []interface{} {
	providers := make([]interface{}, 0)

	providers = append(providers, ProvideIdentityService)

	return providers
}

type IdentityServiceParams struct {
	fx.In

	UserRepo        repository.UserRepository
	Logger          *zap.Logger
	EventDispatcher eventdispatcher.EventDispatcher
	OTPService      otp.OTPService
}

func ProvideIdentityService(params IdentityServiceParams) *IdentityService {
	return &IdentityService{
		userRepo:        params.UserRepo,
		logger:          params.Logger,
		eventDispatcher: params.EventDispatcher,
		otpService:      params.OTPService,
	}
}
