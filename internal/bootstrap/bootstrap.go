package bootstrap

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/nglab-dev/nglab/internal/config"
	"github.com/nglab-dev/nglab/internal/server"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(server.New),
	fx.Provide(config.New),
	// invoke the bootstrap function
	fx.Invoke(bootstrap),
)

func bootstrap(lc fx.Lifecycle, cfg config.Config, srv server.Server) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			// perform any bootstrap actions here
			fmt.Println("Bootstrapping...")

			go func() {
				if err := srv.Start(); err != nil {
					if errors.Is(err, http.ErrServerClosed) {
						slog.Debug("Shutting down the Application")
					} else {
						slog.Error("Error starting the Application: %v", err)
					}
				}
			}()

			return nil
		},
		OnStop: func(context.Context) error {
			// perform any cleanup actions here
			if err := srv.Shutdown(); err != nil {
				return err
			}
			return nil
		},
	})
}
