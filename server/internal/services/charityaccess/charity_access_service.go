package charityaccess

import (
	"context"

	"github.com/google/uuid"
	"github.com/neak-group/nikoogah/internal/app/charity/charity"
	"github.com/neak-group/nikoogah/internal/app/charity/charity/dto"
	"github.com/neak-group/nikoogah/utils/contextutils"
)

type charityAccessServiceImpl struct {
	charityAccessUC *charity.CheckRepresentativeAccessUseCase
}

func (c *charityAccessServiceImpl) CanViewParticipation(ctx context.Context, CharityID uuid.UUID) (bool, error) {
	requesterID, err := contextutils.GetUserIDFromCtx(ctx)
	if err != nil {
		return false, err
	}

	c.charityAccessUC.Execute(ctx, dto.CheckRepresentativeAccessParams{
		CharityID: CharityID,
		UserID:    requesterID,
		AccessKey: "ak.participation.view",
	})
	return false, nil
}

func (c *charityAccessServiceImpl) CanAcceptParticipation(ctx context.Context, CharityID uuid.UUID) (bool, error) {
	requesterID, err := contextutils.GetUserIDFromCtx(ctx)
	if err != nil {
		return false, err
	}

	c.charityAccessUC.Execute(ctx, dto.CheckRepresentativeAccessParams{
		CharityID: CharityID,
		UserID:    requesterID,
		AccessKey: "ak.participation.manage",
	})
	return false, nil
}
