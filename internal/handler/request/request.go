package request

import (
	"github.com/gin-gonic/gin"
	"github.com/nglab-dev/nglab/internal/constant"
	"github.com/nglab-dev/nglab/internal/model"
)

func GetUserClaims(c *gin.Context) (claims *model.UserClaims) {
	claimsValue, exists := c.Get(constant.ClaimsKey)
	if !exists {
		return nil
	}
	claims, exists = claimsValue.(*model.UserClaims)
	if !exists {
		return nil
	}
	return claims
}
