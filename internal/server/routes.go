package server

import (
	"github.com/gin-gonic/gin"
	"github.com/nglab-dev/nglab/internal/handler"
)

func RegisterRoutes(router *gin.Engine) {

	router.Group("")
	{
		router.GET("/login", handler.LoginViewHandler)
		router.POST("/login", handler.LoginHandler)
	}
}
