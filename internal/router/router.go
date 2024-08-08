package router

import (
	"github.com/gin-gonic/gin"
	"github.com/nglab-dev/nglab/internal/db"
	"github.com/nglab-dev/nglab/internal/handler"
	"github.com/nglab-dev/nglab/internal/middleware"
	"github.com/nglab-dev/nglab/internal/service"
	"github.com/nglab-dev/nglab/pkg/env"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {

	jwtSecret := env.GetString("JWT_SECRET", "secret")
	jwtExpire, _ := env.GetInt("JWT_EXPIRE", 3600)

	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// connect db
	db, err := db.Connect()
	if err != nil {
		panic(err)
	}

	// service
	authService := service.NewAuthService(jwtSecret, jwtExpire)
	userService := service.NewUserService(db)

	// handler
	authHandler := handler.NewAuthHandler(authService, userService)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// endpoint
	api := r.Group("/api")
	{
		// no auth
		api.POST("/login", authHandler.Login)

		// auth
		auth := api.Group("").Use(middleware.AuthMiddleware(authService))
		auth.GET("/user", authHandler.GetLoginUser)
	}

	return r
}
