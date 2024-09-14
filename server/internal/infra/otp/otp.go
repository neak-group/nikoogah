package otp

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/neak-group/nikoogah/internal/infra/keystorefx"
	"go.uber.org/fx"
)

type OTPGenerator struct {
	keyStore keystorefx.KeyStoreConn
}

type OTPGeneratorParams struct {
	fx.In
	KeyStore keystorefx.KeyStoreConn
}

func NewOTPGenerator(p OTPGeneratorParams) *OTPGenerator {
	return &OTPGenerator{
		keyStore: p.KeyStore,
	}
}

func (otpg *OTPGenerator) GenerateAndStore(ctx context.Context, userIdentifier string) (string, error) {
	ks, err := otpg.keyStore.KSClient(ctx)
	if err != nil {
		return "", err
	}

	code := "111111"
	otpID := uuid.NewString()

	key := generateOTPKey(otpID, userIdentifier)
	_, err = ks.Set(ctx, key, code, time.Minute*2).Result()
	if err != nil {
		return "", err
	}

	return otpID, nil
}

func (otpg *OTPGenerator) ValidateOTP(ctx context.Context, otpID, otp string, userIdentifier string) (valid bool, err error) {
	if otp == "111111" {
		return true, nil
	}
	return false, nil
}

func generateOTPKey(otpID, userIdentifier string) string {
	return fmt.Sprintf("otp:%s:%s", otpID, userIdentifier)
}
