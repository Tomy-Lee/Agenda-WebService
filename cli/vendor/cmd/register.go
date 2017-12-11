package cmd

import (
	"fmt"
	"os"
	"service"

	"github.com/spf13/cobra"
)

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "register an account",
	Long:  `Register an Agenda account with username, password, email and phone`,
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		email, _ := cmd.Flags().GetString("email")
		phone, _ := cmd.Flags().GetString("phone")
		err := service.Register(username, password, email, phone)
		if err == nil {
			fmt.Println("Registered user:", username)
		} else {
			fmt.Fprintln(os.Stderr, "Error:", err)
		}
	},
}

func init() {
	RootCmd.AddCommand(registerCmd)
	registerCmd.Flags().StringP("username", "u", "", "username")
	registerCmd.Flags().StringP("password", "p", "", "password")
	registerCmd.Flags().StringP("email", "m", "", "your email")
	registerCmd.Flags().StringP("phone", "n", "", "your phone number")
}
