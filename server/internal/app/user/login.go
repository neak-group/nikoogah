package user

import (
	"context"
	"fmt"

	"github.com/neak-group/nikoogah/internal/app/user/dto"
)

func (is *IdentityService) Login(ctx context.Context, input dto.LoginInput) error {
	user, err := is.userRepo.FetchUserByPhone(ctx, input.PhoneNumber)
	if err != nil {
		return err
	}

	//TODO[security]: Record login attempt

	if user == nil {
		return fmt.Errorf("no user found")
	}

	if err = is.otpService.SendOTP(user.PhoneNumber); err != nil {
		return err
	}

	return nil
}

func (is *IdentityService) VerifyLogin(ctx context.Context, input dto.OTPInput) (*dto.UserData, error) {
	user, err := is.userRepo.FetchUserByPhone(ctx, input.PhoneNumber)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, fmt.Errorf("no user found")
	}

	//TODO[Security]: Verify OTP with phone number too
	valid, err := is.otpService.VerifyOTP(input.OTPCode, input.OTPToken)
	if err != nil {
		return nil, err
	}

	if !valid {
		return nil, fmt.Errorf("invalid otp token")
	}

	//TODO[security]: Record login attempt

	return &dto.UserData{
		ID:          user.ID,
		FullName:    fmt.Sprintf(user.FirstName, " ", user.LastName),
		PhoneNumber: user.PhoneNumber,
		UserState:   user.UserState,
	}, nil
}
