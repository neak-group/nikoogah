package participation

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/neak-group/nikoogah/internal/app/rally/rally/dto"
)

// Register a new human participation
func (pc *ParticipationController) RegisterHumanParticipation(c *gin.Context) {
	ctx := c.Request.Context()

	req := new(dto.NewHumanParticipationParams)
	if err := c.Bind(req); err != nil {
		c.Error(err)
		return
	}

	err := pc.newHumanParticipationUseCase.Execute(ctx, req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "Human participation registered successfully",
	})
}
