package user

import (
	"context"
	"fmt"

	"github.com/neak-group/nikoogah/internal/app/user/dto"
	"github.com/neak-group/nikoogah/internal/app/user/entity"
	"github.com/neak-group/nikoogah/internal/core/domain/events"
)

func (is *IdentityService) Verify(ctx context.Context, input *dto.OTPInput) (*dto.UserData, error) {
	user, err := is.userRepo.FetchUserByPhone(ctx, input.PhoneNumber)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, fmt.Errorf("user not found")
	}

	//TODO[Security]: Verify OTP with phone number too
	valid, err := is.otpService.ValidateOTP(ctx, input.OTPToken, input.OTPCode, user.PhoneNumber.PhoneNumber)
	if err != nil {
		return nil, err
	}

	if !valid {
		return nil, fmt.Errorf("invalid otp token")
	}

	if user.UserState == entity.UserPending {
		err = is.userRepo.ChangeUserState(ctx, user.ID, entity.UserActive)
		if err != nil {
			return nil, err
		}
	}

	user.Events = append(user.Events, events.UserJoinedEvent{
		ID:   user.ID,
		Name: fmt.Sprintf("%s %s", user.FirstName, user.LastName),
	})

	if err := is.eventDispatcher.Dispatch(user.Events[0]); err != nil {
		return nil, err
	}

	return &dto.UserData{
		ID:          user.ID,
		FullName:    fmt.Sprintf(user.FirstName, " ", user.LastName),
		PhoneNumber: user.PhoneNumber,
		UserState:   user.UserState,
	}, nil
}
