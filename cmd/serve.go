package cmd

import (
	"io/ioutil"
	"os"

	"github.com/SermoDigital/jose/crypto"

	"github.com/spf13/viper"

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

			keyFile := viper.GetString("security.public_key")
			keyb, err := ioutil.ReadFile(keyFile)
			if err != nil {
				logger.Error(err.Error())
				os.Exit(1)
			}

			key, err := crypto.ParseRSAPublicKeyFromPEM(keyb)
			if err != nil {
				logger.Error(err.Error())
				os.Exit(1)
			}

			server := api.NewServer(logger, key)
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
