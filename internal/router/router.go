package router

import (
	"time"

	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/nglab-dev/nglab/docs"
	"github.com/nglab-dev/nglab/internal/cache"
	"github.com/nglab-dev/nglab/internal/config"
	"github.com/nglab-dev/nglab/internal/db"
	"github.com/nglab-dev/nglab/internal/handler"
	"github.com/nglab-dev/nglab/internal/middleware"
	"github.com/nglab-dev/nglab/internal/service"
	"github.com/nglab-dev/nglab/pkg/log"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter(cfg *config.Config) *gin.Engine {

	// init logger
	log.InitLogger(cfg.Log.Level, cfg.Log.Encoding)

	// create router
	r := gin.New()

	r.Use(ginzap.Ginzap(log.Logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(log.Logger, true))
	r.Use(cors.Default())

	// connect db
	db, err := db.Connect(cfg)
	if err != nil {
		panic(err)
	}

	// init cache
	c, err := cache.Init(cfg)
	if err != nil {
		panic(err)
	}

	// service
	userService := service.NewUserService(db)
	roleService := service.NewRoleService(db)
	authService := service.NewAuthService(cfg.JWT.Secret, cfg.JWT.ExpireTime, db, c, userService)
	dictService := service.NewDictService(db)

	// handler
	authHandler := handler.NewAuthHandler(authService, userService)
	userHandler := handler.NewUserHandler(userService)
	roleHandler := handler.NewRoleHandler(roleService)
	dictHandler := handler.NewDictHandler(dictService)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	docs.SwaggerInfo.BasePath = cfg.Server.Prefix

	// endpoint
	api := r.Group(cfg.Server.Prefix)
	{
		// no auth
		api.POST("/login", authHandler.Login)

		// auth
		auth := api.Group("").Use(middleware.AuthMiddleware(authService))
		auth.POST("/logout", authHandler.Logout)

		auth.GET("/user", userHandler.GetLoginUser)
		auth.PATCH("/user", userHandler.UpdateLoginUser)
		// users
		auth.GET("/users", userHandler.ListUsers)
		auth.POST("/users", userHandler.CreateUser)
		auth.GET("/users/:id", userHandler.GetUser)
		auth.PATCH("/users/:id", userHandler.UpdateUser)
		// roles
		auth.GET("/roles", roleHandler.ListRoles)
		// dicts
		auth.GET("/dicts/types", dictHandler.ListDictTypes)
		auth.POST("/dicts/types", dictHandler.CreateDictType)
		auth.GET("/dicts", dictHandler.ListDicts)
	}

	return r
}
