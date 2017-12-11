package cmd

import (
	"fmt"
	"os"
	"service"

	"github.com/spf13/cobra"
)

var addParticipatorCmd = &cobra.Command{
	Use:   "add-participator",
	Short: "add a participator to a meeting",
	Long:  `add a existed participator to a meeting`,
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		participators, _ := cmd.Flags().GetStringArray("participator")
		err := service.AddParticipatorToMeeting(title, participators)
		if err == nil {
			fmt.Printf("Added participator to the meeting %s\n", title)
		} else {
			fmt.Fprintln(os.Stderr, "Error:", err)
		}
	},
}

func init() {
	RootCmd.AddCommand(addParticipatorCmd)
	addParticipatorCmd.Flags().StringP("title", "t", "", "the title of the meeting")
	addParticipatorCmd.Flags().StringArrayP("participator", "p", nil, "the new participator of the meeting")
}
