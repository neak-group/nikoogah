package volunteer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/neak-group/nikoogah/internal/app/rally/volunteer/dto"
	"go.uber.org/zap"
)

// Fetch the profile of a volunteer by ID
func (vc *VolunteerController) FetchProfile(c *gin.Context) {
	ctx := c.Request.Context()

	req := new(dto.FetchProfileParams)
	if err := c.Bind(req); err != nil {
		c.Error(err)
		return
	}

	profile, err := vc.fetchProfileUseCase.Execute(ctx, req)
	if err != nil {
		vc.logger.Error("Error fetching volunteer profile", zap.Error(err))
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, profile)
}
