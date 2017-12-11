package cmd

import (
	"fmt"
	"os"
	"service"

	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "log out of Agenda",
	Long:  `Log out of Agenda`,
	Run: func(cmd *cobra.Command, args []string) {
		err := service.Logout()
		if err == nil {
			fmt.Println("Logged out successfully")
		} else {
			fmt.Fprintln(os.Stderr, "Error:", err)
		}
	},
}

func init() {
	RootCmd.AddCommand(logoutCmd)
}
