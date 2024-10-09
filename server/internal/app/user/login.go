package user

import (
	"context"
	"fmt"

	"github.com/neak-group/nikoogah/internal/app/user/dto"
)

func (is *IdentityService) Login(ctx context.Context, input *dto.LoginInput) (otpID string, err error) {
	user, err := is.userRepo.FetchUserByPhone(ctx, input.PhoneNumber)
	if err != nil {
		return "", err
	}

	//TODO[security]: Record login attempt

	if user == nil {
		return "", fmt.Errorf("no user found")
	}

	otpID, err = is.otpService.GenerateAndStore(ctx, user.PhoneNumber.PhoneNumber)
	if err != nil {
		return "", err
	}

	return otpID, nil
}
