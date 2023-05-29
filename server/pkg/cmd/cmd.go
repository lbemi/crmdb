package cmd

import (
	"github.com/lbemi/lbemi/cmd/app/option"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/cmd/server"
	"github.com/lbemi/lbemi/pkg/middleware"
	"github.com/lbemi/lbemi/pkg/rctx"
	"github.com/lbemi/lbemi/routes"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

var completedOptions *option.Options

func NewDefaultAppCommand() *cobra.Command {
	var configFile string

	rootCmd := &cobra.Command{
		Use:     "GO-OPS system is for operator.",
		Short:   "GO-OPS system is for operator",
		Example: "go-ops --config config.yaml",
		Version: "1.0.0",
		PreRun: func(cmd *cobra.Command, args []string) {
			completedOptions = option.NewOptions().WithConfig(configFile).WithLog()
			rctx.UserAfterHandlerInterceptor(middleware.LogHandler)
		},
		Run: run,
	}
	rootCmd.Flags().StringVar(&configFile, "config", "", "Set the GO-OPS startup configuration file path")

	return rootCmd
}

func run(cmd *cobra.Command, args []string) {

	httpSever := server.NewHttpSever(":" + completedOptions.Config.App.Port)
	container := httpSever.Container
	container.Filter(middleware.Cors(container).Filter)
	routes.InitTestRouter(container)
	httpSever.Start()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if err := httpSever.Stop(); err != nil {
		log.Logger.Errorf("fault stop server. %s", err)
		os.Exit(-3)
	}
}
