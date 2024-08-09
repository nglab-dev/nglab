package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/nglab-dev/nglab/internal/constant"
	"github.com/nglab-dev/nglab/internal/handler/request"
	"github.com/nglab-dev/nglab/internal/handler/response"
	"github.com/nglab-dev/nglab/internal/service"
)

func AuthMiddleware(authService service.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := request.GetToken(c)

		// validate token
		claims, err := authService.ValidateToken(tokenString)
		if err != nil {
			response.Unauthorized(c, err)
			c.Abort()
		}
		// check token is login

		c.Set(constant.CtxKeyClaims, claims)
		c.Next()
	}
}
