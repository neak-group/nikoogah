package user

import (
	"context"
	"fmt"
	"time"

	"github.com/neak-group/nikoogah/internal/app"
	"github.com/neak-group/nikoogah/internal/app/user/dto"
	"github.com/neak-group/nikoogah/internal/app/user/entity"
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

func (is *IdentityService) RegisterUser(ctx context.Context, input dto.UserInput) error {
	var user *entity.User

	user, err := is.userRepo.FetchUserByPhone(ctx, input.PhoneNumber)
	if err != nil {
		return err
	}

	if user == nil {
		user, err := entity.NewUser(input.FirstName, input.LastName, input.PhoneNumber, input.NationalCode)
		if err != nil {
			return err
		}

		err = is.userRepo.CreateUser(ctx, user)
		if err != nil {
			return err
		}
	} else {
		if user.UserState != entity.UserPending {
			return fmt.Errorf("user already registered")
		}

		user.FirstName = input.FirstName
		user.LastName = input.LastName
		user.NationalCode = input.NationalCode
		user.UpdatedAt = time.Now()
	}

	//TODO[cleanup]: schedule delete after some pending duration

	if err = is.otpService.SendOTP(user.PhoneNumber); err != nil {
		return err
	}

	return nil

}
