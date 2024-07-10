package server

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Port    string `json:"port" yaml:"port" env:"HTTP_PORT" default:"8080"`
	Address string `json:"address" yaml:"address" env:"HTTP_ADDRESS" default:"0.0.0.0"`
}

func Run(config Config, middleware ...gin.HandlerFunc) error {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	// r.Use(middleware)
	// r.Use(ginzap.RecoveryWithZap(Log, true))

	r.GET("/healthz", func(ctx *gin.Context) { ctx.String(200, "OK") })

	srv := &http.Server{
		Addr:    config.Address + ":" + config.Port,
		Handler: r,
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return errors.New("listen: " + err.Error())
	}

	return nil
}
