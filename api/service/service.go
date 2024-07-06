package service

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewAuthService),
	fx.Provide(NewUserService),
	fx.Provide(NewMenuService),
)
