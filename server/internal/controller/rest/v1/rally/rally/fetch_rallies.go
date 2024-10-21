package rally

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (rc *RallyController) FetchRallies(c *gin.Context) {
	ctx := c.Request.Context()

	rallies, err := rc.fetchRalliesUseCase.Execute(ctx)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, rallies)
}
