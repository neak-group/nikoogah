package user

import (
	"context"
	"fmt"
	"time"

	"github.com/neak-group/nikoogah/internal/app/user/dto"
	"github.com/neak-group/nikoogah/internal/app/user/entity"
)

func (is *IdentityService) RegisterUser(ctx context.Context, input *dto.UserInput) (otpID string, err error) {
	var user *entity.User

	user, err = is.userRepo.FetchUserByPhone(ctx, input.PhoneNumber)
	if err != nil {
		// TODO[Clean]: empty token to defined type
		return "", err
	}

	if user == nil {
		user, err = entity.NewUser(input.FirstName, input.LastName, input.PhoneNumber, input.NationalCode)
		if err != nil {
			return "", err
		}

	} else {
		if user.UserState != entity.UserPending {
			return "", fmt.Errorf("user already registered")
		}

		user.FirstName = input.FirstName
		user.LastName = input.LastName
		user.NationalCode = input.NationalCode
		user.UpdatedAt = time.Now()

	}
	err = is.userRepo.SaveUser(ctx, user)
	if err != nil {
		return "", err
	}

	//TODO[cleanup]: schedule delete after some pending duration

	otpID, err = is.otpService.GenerateAndStore(ctx, user.PhoneNumber.PhoneNumber)
	if err != nil {
		return "", err
	}

	return otpID, nil

}
