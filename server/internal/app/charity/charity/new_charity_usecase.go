package charity

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/neak-group/nikoogah/internal/app"
	"github.com/neak-group/nikoogah/internal/app/charity/charity/entity"
	"github.com/neak-group/nikoogah/internal/app/charity/charity/valueobjects"
	"go.uber.org/fx"
)

type RegisterCharityUseCase struct {
	repo CharityRepository
}

type RegisterCharityUCParams struct {
	fx.In

	Repo CharityRepository
}

func ProvideRegisterCharityUC(params RegisterCharityUCParams) *RegisterCharityUseCase {
	return &RegisterCharityUseCase{
		repo: params.Repo,
	}
}

func init() {
	app.RegisterUseCaseProvider(ProvideRegisterCharityUC)
}

type RegisterCharityParams struct {
	Name          string
	Phone         string
	CityPhoneCode string
	Email         string
	Province      string
	City          string
	Address       string
	PostalCode    string
	NationalID    string
	EconomicID    string
	CEO           string
}

func (uc RegisterCharityUseCase) Execute(ctx context.Context, params RegisterCharityParams) (uuid.UUID, error) {
	//TODO: fix aggregate transactions

	tierID, err := uc.repo.FindCharityTierID("basic")
	if err != nil {
		//TODO: fix error
		return uuid.Nil, err
	}

	phone, ok := valueobjects.NewPhone(params.Phone, params.CityPhoneCode)
	if !ok {
		// TODO: fix Error
		return uuid.Nil, fmt.Errorf("invalid phone number")
	}

	address, ok := valueobjects.NewAddress(params.Province, params.City, params.Address, params.PostalCode)
	if !ok {
		// TODO: fix Error
		return uuid.Nil, fmt.Errorf("invalid address")
	}

	email, ok := valueobjects.NewEmail(params.Email)
	if !ok {
		// TODO: fix Error
		return uuid.Nil, fmt.Errorf("invalid email")
	}

	charity := &entity.Charity{
		CharityTierID:   tierID,
		Name:            params.Name,
		Address:         address,
		Phone:           phone,
		EmailAddress:    email,
		NationalID:      params.NationalID,
		EconomicNumber:  params.EconomicID,
		CEO:             params.CEO,
		Representatives: make([]*entity.Representative, 0),
	}

	firstRepresentative := &entity.Representative{
		UserID:   uuid.Nil, // TODO: read user id from context
		Role:     entity.Manager,
		JoinedAt: time.Now(),
	}

	charity.Representatives = append(charity.Representatives, firstRepresentative)

	charityID, err := uc.repo.CreateCharity(charity)
	if err != nil {
		//TODO: fix error
		return uuid.Nil, err
	}

	//TODO: fire event charity created

	return charityID, nil
}
