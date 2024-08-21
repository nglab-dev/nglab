package router

import (
	"log/slog"

	"github.com/nglab-dev/nglab/api/handler"
	"github.com/nglab-dev/nglab/internal/server"
)

type MenuRouter struct {
	srv         server.Server
	menuHandler handler.MenuHandler
}

func NewMenuRouter(srv server.Server, menuHandler handler.MenuHandler) MenuRouter {
	return MenuRouter{
		srv,
		menuHandler,
	}
}

func (r MenuRouter) Setup() {
	api := r.srv.Router
	{
		api.GET("/menus", r.menuHandler.List)
		api.POST("/menus", r.menuHandler.Create)
	}
	slog.Info("Menu router is setup")
}
