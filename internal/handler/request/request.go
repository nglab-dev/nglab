package request

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nglab-dev/nglab/internal/constant"
	"github.com/nglab-dev/nglab/internal/model"
)

func GetUserClaims(c *gin.Context) (claims *model.UserClaims) {
	claimsValue, exists := c.Get(constant.CtxKeyClaims)
	if !exists {
		return nil
	}
	claims, exists = claimsValue.(*model.UserClaims)
	if !exists {
		return nil
	}
	return claims
}

func GetUserID(c *gin.Context) (userID uint) {
	claims := GetUserClaims(c)
	if claims == nil {
		return 0
	}
	return claims.UserID
}

func GetUserName(c *gin.Context) (userName string) {
	claims := GetUserClaims(c)
	if claims == nil {
		return ""
	}
	return claims.Username
}

func GetToken(ctx *gin.Context) (token string) {
	token = ctx.GetHeader(constant.TokenName)
	if len(token) == 0 {
		token = ctx.Query(constant.TokenName)
	}
	return strings.TrimPrefix(token, constant.TokenPrefix)
}
