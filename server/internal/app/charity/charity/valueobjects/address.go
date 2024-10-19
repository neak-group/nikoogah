package valueobjects

import "fmt"

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

func (address Address) String() string {
	return fmt.Sprintf("%s, %s, %s", address.Province, address.City, address.LocalAddress)
}
