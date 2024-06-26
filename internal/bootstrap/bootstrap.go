package bootstrap

import (
	"context"
	"log/slog"

	"github.com/nglab-dev/nglab/api/handler"
	"github.com/nglab-dev/nglab/api/middleware"
	"github.com/nglab-dev/nglab/api/repo"
	"github.com/nglab-dev/nglab/api/router"
	"github.com/nglab-dev/nglab/api/service"
	"github.com/nglab-dev/nglab/internal/config"
	"github.com/nglab-dev/nglab/internal/database"
	"github.com/nglab-dev/nglab/internal/redis"
	"github.com/nglab-dev/nglab/internal/server"
	"go.uber.org/fx"
)

var Module = fx.Options(
	// base modules
	fx.Provide(config.New),
	fx.Provide(server.New),
	fx.Provide(database.New),
	fx.Provide(redis.New),

	// api modules
	handler.Module,
	router.Module,
	middleware.Module,
	service.Module,
	repo.Module,

	// invoke the bootstrap function
	fx.Invoke(bootstrap),
)

func bootstrap(
	lc fx.Lifecycle,
	cfg config.Config,
	db database.Database,
	srv server.Server,
	routes router.Routes,
	middleware middleware.Middlewares,
) {
	dbConn, err := db.DB.DB()
	if err != nil {
		slog.Error("Error connecting to the database: %v", err)
	}
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			// perform any bootstrap actions here
			slog.Info("Starting app...")

			go func() {
				middleware.Setup()
				routes.Setup()

				if err := srv.Start(); err != nil {
					slog.Error("Error starting the Application: %v", err)
				}
			}()

			return nil
		},
		OnStop: func(context.Context) error {
			slog.Info("Stopping app...")
			err := dbConn.Close()
			if err != nil {
				return err
			}
			// perform any cleanup actions here
			if err := srv.Shutdown(); err != nil {
				return err
			}
			return nil
		},
	})
}
