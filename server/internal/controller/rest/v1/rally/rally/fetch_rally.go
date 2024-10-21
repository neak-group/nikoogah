package rally

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/neak-group/nikoogah/internal/app/rally/rally/dto"
)

// Fetch a specific rally by ID
func (rc *RallyController) FetchRally(c *gin.Context) {
	ctx := c.Request.Context()

	req := new(dto.FetchRallyParams)
	if err := c.Bind(req); err != nil {
		c.Error(err)
		return
	}

	rally, err := rc.fetchRallyUseCase.Execute(ctx, req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, rally)
}
