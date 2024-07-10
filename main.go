package main

import (
	"flag"

	"github.com/nglab-dev/nglab/internal/boot"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	configFile := flag.String("c", "config.yaml", "config file path")
	flag.Parse()

	if err := boot.Bootstrap(*configFile); err != nil {
		panic(err)
	}
}
