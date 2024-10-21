package participation

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/neak-group/nikoogah/internal/app/rally/rally/dto"
)

func (pc *ParticipationController) RegisterFundParticipation(c *gin.Context) {
	ctx := c.Request.Context()

	req := new(dto.NewFundParticipationParams)
	if err := c.Bind(req); err != nil {
		c.Error(err)
		return
	}

	err := pc.newFundParticipationUseCase.Execute(ctx, req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "Fund participation registered successfully",
	})
}
