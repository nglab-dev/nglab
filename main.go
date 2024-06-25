package main

import (
	"github.com/nglab-dev/nglab/cmd"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	cmd.Execute()
}
