package main

import (
	"github.com/nglab-dev/nglab/cmd"
)

// @title NGLab API
// @version 1.0
// @description This is a sample server celler server.
// @basePath /api/v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	cmd.Execute()
}
