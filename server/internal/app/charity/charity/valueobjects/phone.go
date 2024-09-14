package valueobjects

type PhoneNumber struct {
	RegionalCode string
	Number       string
}

func (PhoneNumber) PhoneIsValid() bool {
	return true
}

type CellPhoneNumber struct {
	Number string
}

func (CellPhoneNumber) PhoneIsValid() bool {
	return true
}
