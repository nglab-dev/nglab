package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nglab-dev/nglab/internal/config"
	"github.com/nglab-dev/nglab/pkg/log"
)

type HTTPServer struct {
	srv *http.Server
}

func NewHTTPServer(c *config.Config, handler http.Handler) *HTTPServer {
	addr := fmt.Sprintf("%s:%d", c.Server.Host, c.Server.Port)
	return &HTTPServer{
		srv: &http.Server{
			Addr:           addr,
			Handler:        handler,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
	}
}

func (s *HTTPServer) Run() (err error) {
	go func() {
		log.Logger.Sugar().Infof("HTTP server is running on %s", s.srv.Addr)
		if err := s.srv.ListenAndServe(); err != nil {
			log.Logger.Sugar().Errorf("HTTP server error: %v", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// make blocking channel and waiting for a signal
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("error when shutdown server: %v", err)
	}

	// catching ctx.Done(). timeout of 5 seconds.
	<-ctx.Done()
	return
}
