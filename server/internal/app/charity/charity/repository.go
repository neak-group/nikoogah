package charity

import (
	"github.com/google/uuid"
	"github.com/neak-group/nikoogah/internal/app/charity/charity/entity"
)

type CharityRepository interface {
	FindCharityTierID(name string) (uuid.UUID, error)
	CreateCharity(charity *entity.Charity) (uuid.UUID, error)
}
