package charity

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/neak-group/nikoogah/internal/app/charity/charity/dto"
)

func (uc *CharityController) RegisterCharity(c *gin.Context) {
	ctx := c.Request.Context()

	req := new(dto.RegisterCharityParams)
	err := c.Bind(req)
	if err != nil {
		return
	}

	charityID, err := uc.registerUseCase.Execute(ctx, req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "charity registered successfully",
		"data": gin.H{
			"charityId": charityID,
		},
	})
}
