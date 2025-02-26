// cmd/root.go
package cmd

import (
	"github.com/biodigitalJaz/go-cli-terminal/controllers"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "cli-app",
	Short: "An interactive CLI application",
	Long:  `A CLI application with models, controllers, and actions using Cobra and interactive menus.`,
}

func init() {
	// Add subcommands
	RootCmd.AddCommand(controllers.UserCmd)
	RootCmd.AddCommand(controllers.ProjectCmd)
	RootCmd.AddCommand(InteractiveCmd)
}
