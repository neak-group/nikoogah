package valueobjects

type PhoneNumber struct {
	RegionalCode string
	Number       string
}

func NewPhone(phone string, regionalCode string) (PhoneNumber, bool) {
	//TODO: fix validations
	return PhoneNumber{
		RegionalCode: regionalCode,
		Number:       phone,
	}, true
}

type CellPhoneNumber struct {
	Number string
}

func NewCellPhone(phone string) (CellPhoneNumber, bool) {
	//TODO: fix validations
	return CellPhoneNumber{
		Number: phone,
	}, true
}
