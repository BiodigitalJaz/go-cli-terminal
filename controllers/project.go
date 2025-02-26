// controllers/project.go
package controllers

import (
	project "github.com/biodigitalJaz/go-cli-terminal/actions/projects"
	"github.com/spf13/cobra"
)

var ProjectCmd = &cobra.Command{
	Use:   "project",
	Short: "Project management commands",
	Long:  `Commands for managing projects in the application.`,
}

func init() {
	// Add subcommands
	ProjectCmd.AddCommand(&cobra.Command{
		Use:   "list",
		Short: "List all projects",
		Run: func(cmd *cobra.Command, args []string) {
			project.ListProjects()
		},
	})

	ProjectCmd.AddCommand(&cobra.Command{
		Use:   "create",
		Short: "Create a new project",
		Run: func(cmd *cobra.Command, args []string) {
			project.CreateProject()
		},
	})

	ProjectCmd.AddCommand(&cobra.Command{
		Use:   "delete",
		Short: "Delete a project",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			project.DeleteProject(args[0])
		},
	})
}
