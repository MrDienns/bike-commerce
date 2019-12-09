package cmd

import (
	"fmt"

	"github.com/MrDienns/bike-commerce/internal/app/view"
	"github.com/spf13/cobra"
)

var (
	consoleCmd = &cobra.Command{
		Use:   "console",
		Short: "Starts the console application",
		Run: func(cmd *cobra.Command, args []string) {
			screen := view.NewRoot()
			if err := screen.Start(); err != nil {
				fmt.Printf("Failed to start application: %v\n", err)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(consoleCmd)
}
