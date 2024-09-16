package charitytier

import "github.com/neak-group/nikoogah/internal/app/charity/charitytier/entity"

type CharityTierRepository interface {
	SaveCharityTier(*entity.CharityTier) error
}
