package cmd

import (
	"fmt"
	"os"
	"service"

	"github.com/spf13/cobra"
)

var addMeetingCmd = &cobra.Command{
	Use:   "add-meeting",
	Short: "Add a meeting",
	Long:  `Add a meeting with title, participator, start time, end time`,
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		participators, _ := cmd.Flags().GetStringArray("participator")
		startTime, _ := cmd.Flags().GetString("startTime")
		endTime, _ := cmd.Flags().GetString("endTime")
		err := service.AddMeeting(title, participators, startTime, endTime)
		if err == nil {
			fmt.Println("Add meeting:", title)
		} else {
			fmt.Fprintln(os.Stderr, "Error:", err)
		}
	},
}

func init() {
	RootCmd.AddCommand(addMeetingCmd)
	addMeetingCmd.Flags().StringP("title", "t", "", "title of the meeting")
	addMeetingCmd.Flags().StringArrayP("participator", "p", nil, "participators of the meeting")
	addMeetingCmd.Flags().StringP("startTime", "s", "", "start time of the meeting")
	addMeetingCmd.Flags().StringP("endTime", "e", "", "end time of the meeting")
}
