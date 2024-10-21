package participation

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/neak-group/nikoogah/internal/app/rally/rally/dto"
	"github.com/neak-group/nikoogah/utils/uuid"
)

func (pc *ParticipationController) FetchParticipants(c *gin.Context) {
	ctx := c.Request.Context()

	rid, err := uuid.Parse(c.Param("rally-id"))
	if err != nil {
		c.Error(err)
		return
	}

	participants, err := pc.getParticipantsUseCase.Execute(ctx, &dto.GetParticipantsParams{RallyID: rid})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, participants)
}
