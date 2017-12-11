package cmd

import (
	"fmt"
	"os"
	"service"

	"github.com/spf13/cobra"
)

var cancelMeetingCmd = &cobra.Command{
	Use:   "cancel-meeting",
	Short: "cancel a meeting",
	Long:  `cancel a meeting by speecher`,
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		err := service.DeleteFromMeeting(title)
		if err == nil {
			fmt.Printf("Canceled the meeting %s\n", title)
		} else {
			fmt.Fprintln(os.Stderr, "Error:", err)
		}
	},
}

func init() {
	RootCmd.AddCommand(cancelMeetingCmd)
	cancelMeetingCmd.Flags().StringP("title", "t", "", "the title of the meeting")
}
