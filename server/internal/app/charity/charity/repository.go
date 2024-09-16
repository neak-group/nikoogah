package charity

import (
	"github.com/google/uuid"
	"github.com/neak-group/nikoogah/internal/app/charity/charity/entity"
)

type CharityRepository interface {
	FindCharityTierID(name string) (uuid.UUID, error)
	FetchCharity(id uuid.UUID) (*entity.Charity, error)
	CreateCharity(charity *entity.Charity) (uuid.UUID, error)
	SaveCharity(charity *entity.Charity) (uuid.UUID, error)

	FindRepresentativeByUserID(userID uuid.UUID) (*entity.Representative, error)
	FindExistingRepresentativeByUserID(userID uuid.UUID) (bool, error)
}
