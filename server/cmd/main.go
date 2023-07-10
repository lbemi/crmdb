package main

import (
	"os"

	"github.com/lbemi/lbemi/pkg/cmd"
)

func main() {
	command := cmd.NewDefaultAppCommand()
	if err := command.Execute(); err != nil {
		command.PrintErrf("GO-OPS start failed. %v", err)
		os.Exit(1)
	}
}
