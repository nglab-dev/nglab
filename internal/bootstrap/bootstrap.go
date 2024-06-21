package bootstrap

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

	"github.com/nglab-dev/nglab/api/handler"
	"github.com/nglab-dev/nglab/api/repo"
	"github.com/nglab-dev/nglab/api/router"
	"github.com/nglab-dev/nglab/api/service"
	"github.com/nglab-dev/nglab/internal/config"
	"github.com/nglab-dev/nglab/internal/database"
	"github.com/nglab-dev/nglab/internal/serve"
	"go.uber.org/fx"
)

var Module = fx.Options(
	handler.Module,
	router.Module,
	service.Module,
	repo.Module,
	fx.Provide(serve.New),
	fx.Provide(config.New),
	fx.Provide(database.New),
	// invoke the bootstrap function
	fx.Invoke(bootstrap),
)

func bootstrap(lc fx.Lifecycle, cfg config.Config, srv serve.Server, routes router.Routes, db database.Database) {
	dbConn, err := db.ORM.DB()
	if err != nil {
		slog.Error("Error connecting to the database: %v", err)

	}
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			// perform any bootstrap actions here
			slog.Info("Starting the Application")

			go func() {
				routes.Setup()

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
			slog.Info("Stopping the Application")
			// perform any cleanup actions here
			if err := srv.Shutdown(); err != nil {
				return err
			}

			if dbConn != nil {
				return dbConn.Close()
			}
			return nil
		},
	})
}
