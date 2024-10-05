package services

import coreobjects "github.com/neak-group/nikoogah/internal/core/valueobjects"


type OTPService interface{
	SendOTP(PhoneNumber coreobjects.PhoneNumber) (err error)
	VerifyOTP(otpCode string, otpToken string) (verified bool,err error) 
}