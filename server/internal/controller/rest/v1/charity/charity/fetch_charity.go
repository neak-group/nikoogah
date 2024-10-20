package charity

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/neak-group/nikoogah/internal/app/charity/charity/dto"
	"github.com/neak-group/nikoogah/utils/uuid"
)

func (uc *CharityController) FetchCharity(c *gin.Context) {
	ctx := c.Request.Context()

	id, err := uuid.Parse(c.Param("charity-id"))
	if err != nil {
		c.Error(fmt.Errorf("invalid id:%w", err))
		return
	}

	charity, err := uc.fetchCharityUseCase.Execute(ctx, &dto.FetchCharityParams{CharityID: id})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "charity found",
		"data": charity,
	})
}
