package main

import (
	"github.com/nglab-dev/nglab/cmd"
)

// @Title NGLab API
// @Version 1.0
// @Description This is a sample server celler server.
// @BasePath /api/v1
// @SecurityDefinitions.apikey ApiKeyAuth
// @In header
// @Name Authorization
func main() {
	cmd.Execute()
}
