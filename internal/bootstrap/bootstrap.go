package bootstrap

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/nglab-dev/nglab/internal/config"
	"github.com/nglab-dev/nglab/internal/server"
	"github.com/nglab-dev/nglab/internal/storage"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(server.New),
	fx.Provide(config.New),
	fx.Provide(storage.New),
	// invoke the bootstrap function
	fx.Invoke(bootstrap),
)

func bootstrap(lc fx.Lifecycle, srv server.Server, cfg config.Config, st storage.Storage) {
	db, err := st.DB.DB()
	if err != nil {
		slog.Error("Error connecting to the database: %v", err)

	}
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

			if db != nil {
				return db.Close()
			}
			return nil
		},
	})
}
