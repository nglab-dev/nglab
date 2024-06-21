package router

import (
	"github.com/nglab-dev/nglab/api/handler"
	"github.com/nglab-dev/nglab/internal/server"
)

type APIRouter struct {
	srv         server.Server
	authHandler handler.AuthHandler
}

func NewAPIRouter(srv server.Server, authHandler handler.AuthHandler) APIRouter {
	return APIRouter{
		srv,
		authHandler,
	}
}

func (r *APIRouter) Setup() {

	api := r.srv.Router.Group("/api")
	{
		api.GET("/login", r.authHandler.HandleLogin)
	}
}
