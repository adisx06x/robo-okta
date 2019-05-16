package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "robo-okta",
	Short: "Automates cumbersome Okta UI tasks",
	Long:  "\n A collection of automated Okta UI taks that are slow when being done manually.",
}

// Execute executes root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
