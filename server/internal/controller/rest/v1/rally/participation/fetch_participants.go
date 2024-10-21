package participation

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/neak-group/nikoogah/internal/app/rally/rally/dto"
)

func (pc *ParticipationController) FetchParticipants(c *gin.Context) {
	ctx := c.Request.Context()

	req := new(dto.GetParticipantsParams)
	if err := c.Bind(req); err != nil {
		c.Error(err)
		return
	}

	participants, err := pc.getParticipantsUseCase.Execute(ctx, req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, participants)
}
