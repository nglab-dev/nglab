package router

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewViewRouter),
	fx.Provide(NewRouter),
)

// Router is the interface for router.
type IRoute interface {
	Setup()
}

// Routes is the list of routes.
type Routes []IRoute

// New returns a new router.
func NewRouter(viewRouter ViewRouter) Routes {
	return Routes{
		viewRouter,
	}
}

// Setup set all routes.
func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
