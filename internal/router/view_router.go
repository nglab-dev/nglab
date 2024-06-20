package router

import (
	"github.com/nglab-dev/nglab/internal/handler"
	"github.com/nglab-dev/nglab/internal/serve"
)

type ViewRouter struct {
	srv     serve.Server
	handler handler.AuthHandler
}

func NewViewRouter(srv serve.Server, handler handler.AuthHandler) ViewRouter {
	return ViewRouter{
		srv,
		handler,
	}
}

func (r ViewRouter) Setup() {
	api := r.srv.Router
	{
		api.GET("/login", r.handler.HandleLoginView)
		api.GET("/signup", r.handler.HandleSignupView)
	}
}
