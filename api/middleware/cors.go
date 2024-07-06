package middleware

import (
	"log/slog"

	"github.com/gin-contrib/cors"
	"github.com/nglab-dev/nglab/internal/server"
)

// CorsMiddleware configures CORS middleware.
type CorsMiddleware struct {
	srv server.Server
}

// NewCorsMiddleware creates new cors middleware
func NewCorsMiddleware(srv server.Server) CorsMiddleware {
	return CorsMiddleware{
		srv,
	}
}

func (c CorsMiddleware) Setup() {
	c.srv.Engine.Use(cors.Default())
	slog.Info("Cors middleware is setup")
}
