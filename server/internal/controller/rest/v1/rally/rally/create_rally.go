package rally

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/neak-group/nikoogah/internal/app/rally/rally/dto"
)

func (rc *RallyController) CreateRally(c *gin.Context) {
	ctx := c.Request.Context()

	req := new(dto.NewRallyParams)
	if err := c.Bind(req); err != nil {
		c.Error(err)
		return
	}

	rallyID, err := rc.newRallyUseCase.Execute(ctx, req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"rallyId": rallyID,
	})
}
