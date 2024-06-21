package middleware

import (
	"log/slog"
	"time"

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
	cors := cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
	c.srv.Router.Use(cors)
	slog.Info("Cors middleware is setup")
}
