package handler

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewAuthHandler),
	fx.Provide(NewUserHandler),
	fx.Provide(NewMenuHandler),
)
