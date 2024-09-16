package entity

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/neak-group/nikoogah/internal/app/charity/charity/valueobjects"
	"github.com/neak-group/nikoogah/internal/core/service/eventbus"
)

type Charity struct {
	ID             uuid.UUID
	CharityTierID  string
	Name           string
	Address        valueobjects.Address
	Phone          valueobjects.PhoneNumber
	EmailAddress   valueobjects.EmailAddress
	NationalID     string
	EconomicNumber string
	CEO            string

	CreatedAt time.Time
	UpdatedAt time.Time

	Credibility     Credibility
	Representatives []*Representative
	Events          []eventbus.Event
}

func NewCharity(name string) (*Charity, error) {
	charity := &Charity{
		CharityTierID:   "basic",
		Name:            name,
		Representatives: make([]*Representative, 0),
	}

	return charity, nil
}

func (c *Charity) NewAddress(inputProvince, inputCity, inputAddress, inputPostalCode string) error {
	address, ok := valueobjects.NewAddress(inputProvince, inputCity, inputAddress, inputPostalCode)
	if !ok {
		// TODO: fix Error
		return fmt.Errorf("invalid address")
	}

	c.Address = address

	return nil
}

func (c *Charity) NewPhone(inputPhone, inputCityPhoneCode string) error {
	phone, ok := valueobjects.NewPhone(inputPhone, inputCityPhoneCode)
	if !ok {
		// TODO: fix Error
		return fmt.Errorf("invalid phone number")
	}

	c.Phone = phone

	return nil
}

func (c *Charity) NewEmail(inputEmail string) error {
	email, ok := valueobjects.NewEmail(inputEmail)
	if !ok {
		// TODO: fix Error
		return fmt.Errorf("invalid email")
	}

	c.EmailAddress = email

	return nil
}

func (c *Charity) UpdateOfficialData(inputNationalID, inputEconomicalCode, inputCEO string) error {
	//TODO: add validations

	c.CEO = inputCEO
	c.NationalID = inputNationalID
	c.EconomicNumber = inputEconomicalCode

	return nil
}

func (c *Charity) UpdateCharityName(inputName string) error {
	if c.Name == inputName {
		return nil
	}

	// TODO: add validations

	c.Name = inputName
	return nil
}

func (c *Charity) AddRepresentative(userID uuid.UUID, role RepresentativeRole) error {
	if role == "" {
		role = Employee
	}
	c.Representatives = append(c.Representatives, &Representative{
		UserID:   userID, // TODO: read user id from context
		Role:     role,
		JoinedAt: time.Now(),
	})

	return nil
}

func (c *Charity) RemoveRepresentative(userID uuid.UUID, role RepresentativeRole) error {
	for i, r := range c.Representatives {
		if r.UserID == userID{
			c.Representatives = append(c.Representatives[:i], c.Representatives[i+1:]...)
			break
		}
	}

	return nil
}
