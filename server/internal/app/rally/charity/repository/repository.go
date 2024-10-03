package repository

import "github.com/neak-group/nikoogah/internal/app/rally/charity/entity"

type CharityRepository interface {
	SaveCharity(*entity.Charity) error
}
