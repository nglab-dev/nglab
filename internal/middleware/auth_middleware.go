package middleware

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {

	slog.Info("AuthMiddleware")

	c.Next() // 调用后续的处理函数
}
