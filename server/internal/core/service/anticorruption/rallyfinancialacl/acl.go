package rallyfinancialacl

import "context"


type FinancialServiceACL interface{
	RequestPaymentForRallyFee(ctx context.Context, amount string)
}

type FinancialServiceParams struct{
	
}

type FinancialService struct {
	
}

