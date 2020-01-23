package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/SermoDigital/jose/crypto"

	"github.com/spf13/viper"

	"github.com/MrDienns/bike-commerce/pkg/api"

	"github.com/MrDienns/bike-commerce/internal/app/view"
	"github.com/spf13/cobra"
)

var (
	consoleCmd = &cobra.Command{
		Use:   "console",
		Short: "Starts the console application",
		Run: func(cmd *cobra.Command, args []string) {

			publicKeyFile := viper.GetString("security.public_key")
			publickeyb, err := ioutil.ReadFile(publicKeyFile)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
				return
			}

			publickey, err := crypto.ParseRSAPublicKeyFromPEM(publickeyb)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(3)
				return
			}

			screen := view.NewRoot(api.NewClient(publickey, "http://localhost:8080"))
			if err := screen.Start(); err != nil {
				fmt.Printf("Failed to start application: %v\n", err)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(consoleCmd)
}
