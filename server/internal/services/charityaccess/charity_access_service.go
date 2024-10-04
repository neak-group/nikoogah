package charityaccess

type charityAccessServiceImpl struct {
	
}

func (c *charityAccessServiceImpl) CanViewParticipation() (bool, error) {
	return false, nil
}

func (c *charityAccessServiceImpl) CanAcceptParticipation() (bool, error) {
	return false, nil
}
