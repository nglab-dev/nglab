package bootstrap

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nglab-dev/nglab/internal/config"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(config.New),
	// invoke the bootstrap function
	fx.Invoke(bootstrap),
)

func bootstrap(lc fx.Lifecycle, cfg *config.Config) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			// perform any bootstrap actions here
			fmt.Println("Bootstrapping...")

			go func() {
				// TODO: refactor this to use signal pkg
				engine := gin.Default()

				engine.GET("/", func(c *gin.Context) {
					c.JSON(200, cfg)
				})

				engine.Run(cfg.Server.ListenAddr())
			}()

			return nil
		},
		OnStop: func(context.Context) error {
			// perform any cleanup actions here
			return nil
		},
	})
}
