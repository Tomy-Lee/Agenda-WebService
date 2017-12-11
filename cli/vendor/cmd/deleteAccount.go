package cmd

import (
	"fmt"
	"os"
	"service"

	"github.com/spf13/cobra"
)

var deleteAccountCmd = &cobra.Command{
	Use:   "delete-account",
	Short: "delete the current account of Agenda",
	Long:  `delete the current account of Agenda`,
	Run: func(cmd *cobra.Command, args []string) {
		err := service.DeleteUser()
		if err == nil {
			fmt.Println("Deleted Account successfully")
		} else {
			fmt.Fprintln(os.Stderr, "Error:", err)
		}
	},
}

func init() {
	RootCmd.AddCommand(deleteAccountCmd)
}
