package representative

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/neak-group/nikoogah/internal/app/charity/charity/dto"
)

func (rc *RepresentativeController) AddRepresentative(c *gin.Context) {
	ctx := c.Request.Context()

	req := new(dto.AddRepresentativeParams)
	err := c.Bind(req)
	if err != nil {
		return
	}

	err = rc.addRepUseCase.Execute(ctx, req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "representative registered successfully",
	})
}
