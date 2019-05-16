package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of oktaAuto",
	Long:  `All software has versions. This is oktaAuto's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("autoOkta Static Site Generator v0.1 -- HEAD")
	},
}
