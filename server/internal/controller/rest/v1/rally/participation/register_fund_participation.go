package participation

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/neak-group/nikoogah/internal/app/rally/rally/dto"
	"github.com/neak-group/nikoogah/utils/uuid"
)

func (pc *ParticipationController) RegisterFundParticipation(c *gin.Context) {
	ctx := c.Request.Context()

	rid, err := uuid.Parse(c.Param("rally-id"))
	if err != nil {
		c.Error(err)
		return
	}

	req := new(dto.NewFundParticipationParams)
	if err := c.Bind(req); err != nil {
		c.Error(err)
		return
	}
	err = pc.newFundParticipationUseCase.Execute(ctx, &dto.NewFundParticipationParams{
		RallyID:        rid,
		VolunteerID:    req.VolunteerID,
		VolunteerPhone: req.VolunteerPhone,
		Amount:         req.Amount,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "Fund participation registered successfully",
	})
}
