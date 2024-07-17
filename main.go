package main

import (
	"flag"
	"fmt"

	"github.com/nglab-dev/nglab/internal/conf"
)

var configFile = flag.String("c", "config.yaml", "config file path")

func main() {
	flag.Parse()

	config := conf.MustLoad(*configFile)

	fmt.Printf("version: %v\n", config)
}
