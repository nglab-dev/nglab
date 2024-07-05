package middleware

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nglab-dev/nglab/internal/server"
)

// CorsMiddleware configures CORS middleware.
type CorsMiddleware struct {
	srv server.Server
}

// NewCorsMiddleware creates new cors middleware
func NewCorsMiddleware(srv server.Server) CorsMiddleware {
	return CorsMiddleware{
		srv,
	}
}

func (c CorsMiddleware) core() gin.HandlerFunc {

	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

func (c CorsMiddleware) Setup() {
	c.srv.Router.Use(c.core())
	slog.Info("Cors middleware is setup")
}
