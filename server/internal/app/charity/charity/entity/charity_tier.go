package entity

type CharityTier struct {
	name                string
	representativeLimit int
	rallyLimit          int
}

func (ct CharityTier) GetRepresentativeLimit() int {
	return ct.representativeLimit
}

func (ct CharityTier) GetRallyLimit() int {
	return ct.rallyLimit
}

var TierMap = map[string]*CharityTier{
	"basic": Basictier,
}

var Basictier = &CharityTier{
	name:                "basic-tier",
	representativeLimit: 5,
	rallyLimit:          20,
}
