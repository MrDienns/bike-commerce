package cmd

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "bike-commerce",
		Short: "Root command for using the bike-commerce application",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
)

// Execute is the main entry point for the root command. This function will invoke the Cobra implementation of the
// command, which, in this case, will output the usage guides.
func Execute() {
	rootCmd.Execute()
}

// init gets called by cobra before the command is executed.
func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")
}

// initConfig tries to search for a suitable configuration file and will load the values into environment variables.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath("/etc/fontys/bike-commerce")
		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigName("bike-commerce")
	}
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
