package valueobjects

type EmailAddress string

func (e EmailAddress) EmailIsValid() (valid bool) {
	return true
}
