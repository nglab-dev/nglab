package repo

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewUserRepo),
	fx.Provide(NewMenuRepo),
)
