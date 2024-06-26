package router

import (
	"log/slog"

	"github.com/nglab-dev/nglab/api/handler"
	"github.com/nglab-dev/nglab/internal/server"
)

type AuthRouter struct {
	srv         server.Server
	authHandler handler.AuthHandler
}

func NewAuthRouter(srv server.Server, authHandler handler.AuthHandler) AuthRouter {
	return AuthRouter{
		srv,
		authHandler,
	}
}

func (r AuthRouter) Setup() {
	api := r.srv.Router
	{
		api.POST("/login", r.authHandler.Login)
		api.POST("/register", r.authHandler.Register)
		api.GET("/user", r.authHandler.GetUser)
	}
	slog.Info("Auth router is setup")
}
