package cmd

import (
	"fmt"
	"os"
	"service"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "show registered users",
	Long:  `show registered users`,
	Run: func(cmd *cobra.Command, args []string) {
		results, err := service.FindAll()
		if err == nil {
			fmt.Println("All registered users")
			for i, user := range results {
				fmt.Printf("No. %d\n", i)
				fmt.Printf("Username: %s\n", user.Username)
				fmt.Printf("Email: %s\n", user.Email)
				fmt.Printf("Phone: %s\n\n", user.Phone)
			}
		} else {
			fmt.Fprintln(os.Stderr, "Error:", err)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
