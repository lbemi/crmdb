package main

import (
	"github.com/lbemi/lbemi/pkg/cmd"
	"os"
)

func main() {
	command := cmd.NewDefaultAppCommand()
	if err := command.Execute(); err != nil {
		command.PrintErrf("GO-OPS start failed. %v", err)
		os.Exit(1)
	}
}
