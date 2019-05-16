package cmd

import (
	"fmt"

	"github.com/adisx06x/robo-okta/api"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(usersToGroupCmd)
}

var usersToGroupCmd = &cobra.Command{
	Use:   "add-users-to-group",
	Short: "Add list of users to group in Okta",
	Long:  "\n With add-users-to-group you can add a list of users from a CSV and add them to a Group using a Group ID.",
	Run: func(cmd *cobra.Command, args []string) {
		client := api.NewOktaClient()

		finalArray := api.ReadCSV(filepath)
		Ids := client.GetUserIDs(finalArray)

		client.AddIDsToGroup(Ids, groupID)

		fmt.Printf("Added, users %s to groupID %s.\n", finalArray, groupID)
	},
}

func init() {
	usersToGroupCmd.Flags().StringVarP(&filepath, "filepath", "f", "", "path to csv filename of user emails")
	usersToGroupCmd.MarkFlagRequired("filepath")
	usersToGroupCmd.Flags().StringVarP(&groupID, "groupid", "g", "", "Okta group ID to add users to")
	usersToGroupCmd.MarkFlagRequired("groupID")
}
