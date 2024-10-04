package charityaccess

import (
	"github.com/neak-group/nikoogah/internal/app/rally/rally/services"
	"go.uber.org/fx"
)

type CharityAccessServiceParams struct {
	fx.In
}

type CharityAccessServiceResult struct {
	fx.Out
	
	Rally services.CharityAccessService
}

func ProvideCharityAccessService(params CharityAccessServiceParams) CharityAccessServiceResult{
	impl := &charityAccessServiceImpl{
		
	}
	
	return CharityAccessServiceResult{
		
	}
}
