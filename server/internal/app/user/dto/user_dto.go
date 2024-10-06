package dto

type UserInput struct {
	FirstName    string
	LastName     string
	PhoneNumber  string
	NationalCode string
}

type OTPInput struct {
	PhoneNumber string
	OTPCode     string
	OTPToken    string
}


type LoginInput struct{
	PhoneNumber string
}