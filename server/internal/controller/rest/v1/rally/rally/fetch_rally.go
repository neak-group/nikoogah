package rally

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/neak-group/nikoogah/internal/app/rally/rally/dto"
	"github.com/neak-group/nikoogah/utils/uuid"
)

// Fetch a specific rally by ID
func (rc *RallyController) FetchRally(c *gin.Context) {
	ctx := c.Request.Context()

	rid, err := uuid.Parse(c.Param("rally-id"))
	if err != nil {
		c.Error(err)
		return
	}

	rally, err := rc.fetchRallyUseCase.Execute(ctx, &dto.FetchRallyParams{RallyID: rid})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, rally)
}
