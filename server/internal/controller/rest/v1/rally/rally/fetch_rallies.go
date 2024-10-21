package rally

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/neak-group/nikoogah/internal/app/rally/rally/dto"
	"github.com/neak-group/nikoogah/utils/uuid"
	"go.uber.org/zap"
)

// Fetch all rallies or rallies for a specific charity based on query parameters
func (rc *RallyController) FetchRallies(c *gin.Context) {
	ctx := c.Request.Context()

	// Check if query parameters are provided
	charityIDParam := c.Query("charityId")
	onlyActiveParam := c.Query("onlyActive")

	// If charityId is provided, fetch rallies for the specific charity
	if charityIDParam != "" {
		charityID, err := uuid.Parse(charityIDParam)
		if err != nil {
			rc.logger.Error("Invalid charity ID", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid charity ID"})
			return
		}

		// Check for the onlyActive query parameter (optional)
		onlyActive := onlyActiveParam == "true"

		// Create params object for fetching charity rallies
		req := &dto.FetchCharityRalliesParams{
			CharityID:  charityID,
			OnlyActive: onlyActive,
		}

		// Fetch the rallies for the specific charity
		rallies, err := rc.fetchCharityRalliesUseCase.Execute(ctx, req)
		if err != nil {
			rc.logger.Error("Error fetching charity rallies", zap.Error(err))
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, rallies)
		return
	}

	// If no charityId is provided, fetch all rallies
	rallies, err := rc.fetchRalliesUseCase.Execute(ctx)
	if err != nil {
		rc.logger.Error("Error fetching rallies", zap.Error(err))
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, rallies)
}
