package main

import (
	_ "github.com/lbemi/lbemi/docs"
	"github.com/lbemi/lbemi/pkg/cmd"
	"os"
)

// @title lbemi API 文档
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host 127.0.0.1:8080
// @BasePath /
func main() {
	//server.Run()

	command := cmd.NewDefaultAppCommand()
	if err := command.Execute(); err != nil {
		command.PrintErrf("GO-OPS start failed. %", err)
		os.Exit(1)
	}
}
