package cmd

import (
	"os"

	"github.com/MrDienns/bike-commerce/pkg/api"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Starts the API",
		Run: func(cmd *cobra.Command, args []string) {

			logger, _ := zap.NewDevelopment()
			defer logger.Sync()

			server := api.NewServer(logger)
			if err := server.Start(); err != nil {
				logger.Error(err.Error())
				os.Exit(1)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(serveCmd)
}
