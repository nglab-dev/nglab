package router

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewAuthRouter),
	fx.Provide(NewUserRouter),
	fx.Provide(NewMenuRouter),
	fx.Provide(NewRouter),
)

// Router is the interface for router.
type IRoute interface {
	Setup()
}

// Routes is the list of routes.
type Routes []IRoute

// New returns a new router.
func NewRouter(authRouter AuthRouter, userRouter UserRouter, menuRouter MenuRouter) Routes {
	return Routes{
		authRouter,
		userRouter,
		menuRouter,
	}
}

// Setup set all routes.
func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
