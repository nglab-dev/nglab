package router

import (
	"log/slog"

	"github.com/nglab-dev/nglab/api/handler"
	"github.com/nglab-dev/nglab/internal/server"
)

type UserRouter struct {
	srv         server.Server
	userHandler handler.UserHandler
}

func NewUserRouter(srv server.Server, userHandler handler.UserHandler) UserRouter {
	return UserRouter{
		srv,
		userHandler,
	}
}

func (r UserRouter) Setup() {
	api := r.srv.Router
	{
		api.GET("/users/:id", r.userHandler.Get)
	}
	slog.Info("User router is setup")
}
