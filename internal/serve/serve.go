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
	"github.com/nglab-dev/nglab/public"
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

	render := router.HTMLRender
	router.HTMLRender = &HTMLTemplRenderer{FallbackHtmlRenderer: render}

	// Disable trusted proxy warning.
	router.SetTrustedProxies(nil)

	if cfg.App.IsDev() {
		// disable caching
		router.Use(func(c *gin.Context) {
			c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
			c.Header("Pragma", "no-cache")
			c.Header("Expires", "0")
			c.Next()
		})
		router.StaticFS("/public", http.Dir("./public"))
	} else {
		router.StaticFS("/public", http.FS(public.AssetsFS))
	}

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
