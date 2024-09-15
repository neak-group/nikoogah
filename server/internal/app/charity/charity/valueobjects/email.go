package valueobjects

type EmailAddress string

func NewEmail(email string) (EmailAddress, bool) {

	//TODO: fix validation

	return EmailAddress(email), true
}
