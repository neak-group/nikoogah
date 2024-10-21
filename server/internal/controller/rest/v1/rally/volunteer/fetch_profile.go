package volunteer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Fetch the profile of a volunteer by ID
func (vc *VolunteerController) FetchMyProfile(c *gin.Context) {
	ctx := c.Request.Context()

	profile, err := vc.fetchProfileUseCase.Execute(ctx)
	if err != nil {
		vc.logger.Error("Error fetching volunteer profile", zap.Error(err))
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, profile)
}
