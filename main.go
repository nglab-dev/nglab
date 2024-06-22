package main

import (
	"github.com/nglab-dev/nglab/cmd"
)

// @basePath /api/v1
// @securityDefinitions.apikey Authorization
// @in header
// @name Authorization
func main() {
	cmd.Execute()
}
