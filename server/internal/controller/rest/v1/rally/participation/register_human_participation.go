package participation

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/neak-group/nikoogah/internal/app/rally/rally/dto"
	"github.com/neak-group/nikoogah/utils/uuid"
)

// Register a new human participation
func (pc *ParticipationController) RegisterHumanParticipation(c *gin.Context) {
	ctx := c.Request.Context()

	rid, err := uuid.Parse(c.Param("rally-id"))
	if err != nil {
		c.Error(err)
		return
	}

	req := new(dto.NewHumanParticipationParams)
	if err := c.Bind(req); err != nil {
		c.Error(err)
		return
	}

	err = pc.newHumanParticipationUseCase.Execute(ctx, &dto.NewHumanParticipationParams{
		RallyID:         rid,
		VolunteerID:     req.VolunteerID,
		VolunteerPhone:  req.VolunteerPhone,
		VolunteerEmail:  req.VolunteerEmail,
		VolunteerResume: req.VolunteerResume,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "Human participation registered successfully",
	})
}
