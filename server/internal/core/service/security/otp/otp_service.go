package services

import "context"

type OTPService interface {
	GenerateAndStore(ctx context.Context, userIdentifier string) (string, error)
	ValidateOTP(ctx context.Context, otpID string, otp string, userIdentifier string) (valid bool, err error)
}
