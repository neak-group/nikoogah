package entity

type CharityTier struct {
	name                string
	representativeLimit int
}

func (ct CharityTier) GetRepresentativeLimit() int {
	return ct.representativeLimit
}

var TierMap = map[string]CharityTier{
	"basic": Basictier,
}

var Basictier = CharityTier{
	name:                "basic-tier",
	representativeLimit: 5,
}
