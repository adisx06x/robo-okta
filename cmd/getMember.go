package cmd

import (
	"fmt"
	"log"

	"github.com/adisx06x/robo-okta/api"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getMemberCmd)
}

var getMemberCmd = &cobra.Command{
	Use:   "get-user [login]",
	Short: "Fetches a specific user when you know the user's login",
	Long:  "\n With get-user you can fetch a specific user with their email.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client := api.NewOktaClient()

		userLogin := args[0]
		user, _, err := client.User.GetUser(userLogin)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%+v\n", user)

	},
}
