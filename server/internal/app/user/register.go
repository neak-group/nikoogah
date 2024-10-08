package user

import (
	"context"
	"fmt"
	"time"

	"github.com/neak-group/nikoogah/internal/app/user/dto"
	"github.com/neak-group/nikoogah/internal/app/user/entity"
)

func (is *IdentityService) RegisterUser(ctx context.Context, input dto.UserInput) error {
	var user *entity.User

	user, err := is.userRepo.FetchUserByPhone(ctx, input.PhoneNumber)
	if err != nil {
		return err
	}

	if user == nil {
		user, err = entity.NewUser(input.FirstName, input.LastName, input.PhoneNumber, input.NationalCode)
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
	err = is.userRepo.SaveUser(ctx, user)
	if err != nil {
		return err
	}

	//TODO[cleanup]: schedule delete after some pending duration

	if err = is.otpService.SendOTP(user.PhoneNumber); err != nil {
		return err
	}

	return nil

}

func (is *IdentityService) VerifyRegistration(ctx context.Context, input dto.OTPInput) error {
	user, err := is.userRepo.FetchUserByPhone(ctx, input.PhoneNumber)
	if err != nil {
		return err
	}

	if user == nil {
		return fmt.Errorf("user not found")
	}

	//TODO[Security]: Verify OTP with phone number too
	valid, err := is.otpService.VerifyOTP(input.OTPCode, input.OTPToken)
	if err != nil {
		return err
	}

	if !valid {
		return fmt.Errorf("invalid otp token")
	}

	err = is.userRepo.ChangeUserState(ctx, user.ID, entity.UserActive)
	if err != nil {
		return err
	}

	return nil
}
