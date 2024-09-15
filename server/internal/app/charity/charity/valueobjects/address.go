package valueobjects

type Address struct {
	Province     string
	City         string
	LocalAddress string
	PostalCode   string
}

func NewAddress(province, city, local, postalcode string) (Address, bool) {
	//TODO: fix validation
	return Address{
		Province:     province,
		City:         city,
		LocalAddress: local,
		PostalCode:   postalcode,
	}, true
}
