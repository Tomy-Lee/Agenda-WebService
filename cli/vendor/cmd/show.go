package cmd

import (
	"fmt"
	"os"
	"service"

	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "show meetings",
	Long:  `Show meetings between the time`,
	Run: func(cmd *cobra.Command, args []string) {
		startTime, _ := cmd.Flags().GetString("startTime")
		endTime, _ := cmd.Flags().GetString("endTime")
		results, err := service.FindMeetingByTime(startTime, endTime)
		if err == nil {
			if len(results) == 0 {
				fmt.Printf("No meetings found")
				return
			}
			for i, meeting := range results {
				fmt.Printf("No. %d\n", i)
				fmt.Printf("Title: %s\n", meeting.Title)
				fmt.Printf("Speecher: %s\n", meeting.Speecher)
				fmt.Printf("Participators:\n")
				for _, participator := range meeting.Participators {
					fmt.Printf("             %s\n", participator)
				}
				fmt.Printf("Start Time: %s\n", meeting.StartTime)
				fmt.Printf("End Time: %s\n\n", meeting.EndTime)
			}
		} else {
			fmt.Fprintln(os.Stderr, "Error:", err)
		}
	},
}

func init() {
	RootCmd.AddCommand(showCmd)
	showCmd.Flags().StringP("startTime", "s", "", "start time for query")
	showCmd.Flags().StringP("endTime", "e", "", "end time for query")
}
