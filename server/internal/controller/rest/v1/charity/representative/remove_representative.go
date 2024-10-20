package representative

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/neak-group/nikoogah/internal/app/charity/charity/dto"
	"github.com/neak-group/nikoogah/utils/uuid"
)

func (rc *RepresentativeController) RemoveRepresentative(c *gin.Context) {
	ctx := c.Request.Context()

	charityID, err := uuid.Parse(c.Param("charity-id"))
	if err != nil {
		c.Error(err)
		return
	}

	repID, err := uuid.Parse(c.Param("rep-id"))
	if err != nil {
		c.Error(err)
		return
	}

	err = rc.addRepUseCase.Execute(ctx, &dto.AddRepresentativeParams{
		CharityID: charityID,
		UserID:    repID,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "representative removed successfully",
	})
}
