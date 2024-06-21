package middleware

import (
	"log/slog"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nglab-dev/nglab/api/service"
	"github.com/nglab-dev/nglab/internal/config"
	"github.com/nglab-dev/nglab/internal/server"
)

var _ IMiddleware = (*AuthMiddleware)(nil)

type AuthMiddleware struct {
	cfg         config.Config
	srv         server.Server
	authService service.AuthService
}

func NewAuthMiddleware(cfg config.Config, srv server.Server, authService service.AuthService) AuthMiddleware {
	return AuthMiddleware{
		cfg,
		srv,
		authService,
	}
}

func (auth AuthMiddleware) core() gin.HandlerFunc {
	ignorePaths := auth.cfg.Auth.IgnorePaths

	return func(ctx *gin.Context) {
		if isIgnorePath(ctx.Request.URL.Path, ignorePaths...) {
			ctx.Next()
			return
		}

		tokenString := ExtractToken(ctx)
		claims, err := auth.authService.ValidateToken(tokenString)
		if err != nil {

			ctx.Abort()
		}
		ctx.Set("user", claims)
		ctx.Next()
	}
}

// ExtractToken extracts the token from the request Header.
func ExtractToken(ctx *gin.Context) (token string) {
	token = ctx.GetHeader("Authorization")
	if len(token) == 0 {
		token = ctx.Query("Authorization")
	}
	return strings.TrimPrefix(token, "Bearer ")
}

func (auth AuthMiddleware) Setup() {
	if !auth.cfg.Auth.Enabled {
		return
	}
	auth.srv.Router.Use(auth.core())
	slog.Info("Auth middleware is setup")
}
