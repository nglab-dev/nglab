package server

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nglab-dev/nglab/internal/config"
)

const DefaultShutdownTimeout = time.Minute

var _ IServer = (*Server)(nil)

type Server struct {
	srv *http.Server
}

func New(conf config.Config) Server {
	engine := gin.New()

	engine.GET("/", func(c *gin.Context) {
		c.JSON(200, conf)
	})

	srv := &http.Server{
		Addr:    conf.Server.ListenAddr(),
		Handler: engine,
	}

	return Server{
		srv,
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
