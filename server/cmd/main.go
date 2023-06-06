package main

import (
	_ "github.com/lbemi/lbemi/docs"
	"github.com/lbemi/lbemi/pkg/cmd"
	"os"
)

func main() {
	//server.Run()

	command := cmd.NewDefaultAppCommand()
	if err := command.Execute(); err != nil {
		command.PrintErrf("GO-OPS start failed. %", err)
		os.Exit(1)
	}
}
