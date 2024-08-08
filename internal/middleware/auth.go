package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nglab-dev/nglab/internal/constant"
	"github.com/nglab-dev/nglab/internal/handler/response"
	"github.com/nglab-dev/nglab/internal/service"
)

func AuthMiddleware(authService service.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := ExtractToken(c)
		claims, err := authService.ValidateToken(tokenString)
		if err != nil {
			response.Unauthorized(c, err)
			c.Abort()
		}
		c.Set(constant.ClaimsKey, claims)
		c.Next()
	}
}

func ExtractToken(ctx *gin.Context) (token string) {
	token = ctx.GetHeader(constant.TokenName)
	if len(token) == 0 {
		token = ctx.Query(constant.TokenName)
	}
	return strings.TrimPrefix(token, constant.TokenPrefix)
}
