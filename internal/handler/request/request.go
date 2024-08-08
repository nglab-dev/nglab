package request

import (
	"github.com/gin-gonic/gin"
	"github.com/nglab-dev/nglab/internal/constant"
	"github.com/nglab-dev/nglab/internal/model"
)

func GetLoginUser(c *gin.Context) (user *model.User) {
	claims, exists := c.Get(constant.ClaimsKey)
	if !exists {
		return nil
	}
	user, exists = claims.(*model.User)
	if !exists {
		return nil
	}
	return user
}
