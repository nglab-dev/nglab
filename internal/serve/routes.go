package serve

import (
	"github.com/gin-gonic/gin"
	"github.com/nglab-dev/nglab/internal/handler"
	"github.com/nglab-dev/nglab/internal/middleware"
)

func RegisterRoutes(router *gin.Engine) {

	router.GET("/login", handler.HandleLoginView)
	router.GET("/signup", handler.HandleSignupView)

	router.Use(middleware.AuthMiddleware)
	{
		router.POST("/login", handler.HandleLogin)
		router.POST("/signup", handler.HandleSignup)
	}
}
