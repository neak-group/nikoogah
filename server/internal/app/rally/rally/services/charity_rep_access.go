package services

type CharityAccessService interface {
	CanViewParticipation() (bool, error)
	CanAcceptParticipation() (bool, error)
}

