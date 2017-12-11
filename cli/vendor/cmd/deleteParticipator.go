package cmd

import (
	"fmt"
	"os"
	"service"

	"github.com/spf13/cobra"
)

var deleteParticipatorCmd = &cobra.Command{
	Use:   "delete-participator",
	Short: "delete a participator from a meeting",
	Long:  `delete a existed participator from a meeting`,
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		participators, _ := cmd.Flags().GetStringArray("participator")
		err := service.DeleteParticipatorFromMeeting(title, participators)
		if err == nil {
			fmt.Printf("Deleted participator from the meeting %s\n", title)
		} else {
			fmt.Fprintln(os.Stderr, "Error:", err)
		}
	},
}

func init() {
	RootCmd.AddCommand(deleteParticipatorCmd)
	deleteParticipatorCmd.Flags().StringP("title", "t", "", "the title of the meeting")
	deleteParticipatorCmd.Flags().StringArrayP("participator", "p", nil, "the participator of the meeting")
}
