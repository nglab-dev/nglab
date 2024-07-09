package server

import (
	"errors"
	"net/http"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/nglab-dev/nglab/internal/conf"
	"github.com/nglab-dev/nglab/internal/log"
)

func Run() error {
	config := conf.Get().Server
	logger := log.Get()

	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))

	r.GET("/healthz", func(ctx *gin.Context) { ctx.String(200, "OK") })

	srv := &http.Server{
		Addr:    config.Address + ":" + config.Port,
		Handler: r,
	}

	logger.Info("Starting server on " + config.Address + ":" + config.Port)

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return errors.New("listen: " + err.Error())
	}
	return nil
}
