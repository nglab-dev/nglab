package router

import (
	"github.com/gin-gonic/gin"
	"github.com/nglab-dev/nglab/internal/handler"
)

func RegisterRoutes(r *gin.Engine) {

	api := r.Group("/api")
	{
		api.POST("/login", handler.Login)
		api.POST("/register", handler.Register)
	}
}
