package cmd

import (
	"io/ioutil"
	"os"

	"github.com/MrDienns/bike-commerce/pkg/database"

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

			publicKeyFile := viper.GetString("security.public_key")
			publickeyb, err := ioutil.ReadFile(publicKeyFile)
			if err != nil {
				logger.Error(err.Error())
				os.Exit(1)
				return
			}

			privateKeyFile := viper.GetString("security.private_key")
			privatekeyb, err := ioutil.ReadFile(privateKeyFile)
			if err != nil {
				logger.Error(err.Error())
				os.Exit(2)
				return
			}

			publickey, err := crypto.ParseRSAPublicKeyFromPEM(publickeyb)
			if err != nil {
				logger.Error(err.Error())
				os.Exit(3)
				return
			}

			privatekey, err := crypto.ParseRSAPrivateKeyFromPEM(privatekeyb)
			if err != nil {
				logger.Error(err.Error())
				os.Exit(4)
				return
			}

			connector := database.MySQL{
				Username: viper.GetString("database.username"),
				Password: viper.GetString("database.password"),
				Host:     viper.GetString("database.host"),
				Port:     viper.GetInt("database.port"),
				Database: viper.GetString("database.database"),
			}

			err = connector.Connect()
			if err != nil {
				logger.Error(err.Error())
				os.Exit(5)
				return
			}

			var databaseConnector database.Connector
			databaseConnector = &connector

			server := api.NewServer(logger, publickey, privatekey, databaseConnector)
			if err := server.Start(); err != nil {
				logger.Error(err.Error())
				os.Exit(6)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(serveCmd)
}
