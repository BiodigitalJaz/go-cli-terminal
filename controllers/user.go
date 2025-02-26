// controllers/user.go
package controllers

import (
	"github.com/biodigitalJaz/go-cli-terminal/actions/user"
	"github.com/spf13/cobra"
)

var UserCmd = &cobra.Command{
	Use:   "user",
	Short: "User management commands",
	Long:  `Commands for managing users in the application.`,
}

func init() {
	// Add subcommands
	UserCmd.AddCommand(&cobra.Command{
		Use:   "list",
		Short: "List all users",
		Run: func(cmd *cobra.Command, args []string) {
			user.ListUsers()
		},
	})

	UserCmd.AddCommand(&cobra.Command{
		Use:   "create",
		Short: "Create a new user",
		Run: func(cmd *cobra.Command, args []string) {
			user.CreateUser()
		},
	})

	UserCmd.AddCommand(&cobra.Command{
		Use:   "delete",
		Short: "Delete a user",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			user.DeleteUser(args[0])
		},
	})
}
