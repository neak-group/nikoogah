package rally // Fetch rallies for a specific charity
import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/neak-group/nikoogah/internal/app/rally/rally/dto"
)

func (rc *RallyController) FetchCharityRallies(c *gin.Context) {
	ctx := c.Request.Context()

	req := new(dto.FetchCharityRalliesParams)
	if err := c.Bind(req); err != nil {
		c.Error(err)
		return
	}

	rallies, err := rc.fetchCharityRalliesUseCase.Execute(ctx, req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, rallies)
}
