package serve

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nglab-dev/nglab/internal/config"
	sloggin "github.com/samber/slog-gin"
)

const DefaultShutdownTimeout = time.Minute

type Server struct {
	srv    *http.Server
	Router *gin.Engine
}

func New(cfg config.Config) Server {

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	router := gin.New()

	router.Use(sloggin.New(logger))
	router.Use(gin.Recovery())

	// Disable trusted proxy warning.
	router.SetTrustedProxies(nil)

	srv := &http.Server{
		Addr:    cfg.Server.ListenAddr(),
		Handler: router,
	}

	return Server{
		srv,
		router,
	}
}

func (s *Server) Start() (err error) {
	err = s.srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultShutdownTimeout)
	defer cancel()
	return s.srv.Shutdown(ctx)
}
