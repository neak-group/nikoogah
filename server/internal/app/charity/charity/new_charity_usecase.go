package charity

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/neak-group/nikoogah/internal/app"
	"github.com/neak-group/nikoogah/internal/app/charity/charity/entity"
	"github.com/neak-group/nikoogah/internal/core/service/eventbus"
	"go.uber.org/fx"
)

type RegisterCharityUseCase struct {
	repo CharityRepository
}

type RegisterCharityUCParams struct {
	fx.In

	Repo     CharityRepository
	EventBus eventbus.EventBus
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

	charity, err := entity.NewCharity(params.Name)
	if err != nil {
		return uuid.Nil, err
	}

	if err := charity.NewAddress(params.Province, params.City, params.Address, params.PostalCode); err != nil {
		return uuid.Nil, err
	}

	if err := charity.NewPhone(params.Phone, params.CityPhoneCode); err != nil {
		return uuid.Nil, err
	}

	if err := charity.NewEmail(params.Email); err != nil {
		return uuid.Nil, err
	}

	if err := charity.UpdateOfficialData(params.NationalID, params.EconomicID, params.CEO); err != nil {
		return uuid.Nil, err
	}

	firstRepresentative := &entity.Representative{
		UserID:   uuid.Nil, // TODO: read user id from context
		Role:     entity.Manager,
		JoinedAt: time.Now(),
	}
	if err != nil {
		return uuid.Nil, err
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
