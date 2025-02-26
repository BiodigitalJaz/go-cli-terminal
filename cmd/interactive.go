// cmd/interactive.go
package cmd

import (
	"fmt"
	"os"

	project "github.com/biodigitalJaz/go-cli-terminal/actions/projects"
	"github.com/biodigitalJaz/go-cli-terminal/actions/user"
	"github.com/biodigitalJaz/go-cli-terminal/models"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var InteractiveCmd = &cobra.Command{
	Use:   "interactive",
	Short: "Start the interactive menu",
	Run: func(cmd *cobra.Command, args []string) {
		runInteractiveMenu()
	},
}

func runInteractiveMenu() {
	mainMenu := models.NewOptionList()

	// Add main menu options
	mainMenu.AddOption("User Management", "Manage application users", runUserMenu)
	mainMenu.AddOption("Project Management", "Manage application projects", runProjectMenu)
	mainMenu.AddOption("Exit", "Exit the application", func() error {
		fmt.Println("Goodbye!")
		os.Exit(0)
		return nil
	})

	for {
		prompt := promptui.Select{
			Label: "Select an option",
			Items: mainMenu.GetOptionNames(),
		}

		index, _, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed: %v\n", err)
			return
		}

		option := mainMenu.GetOption(index)
		if option != nil {
			if err := option.Action(); err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		}
	}
}

func runUserMenu() error {
	userMenu := models.NewOptionList()

	// Add user menu options
	userMenu.AddOption("List Users", "List all users", user.ListUsers)
	userMenu.AddOption("Create User", "Create a new user", user.CreateUser)
	userMenu.AddOption("Delete User", "Delete an existing user", promptDeleteUser)
	userMenu.AddOption("Back", "Return to main menu", func() error {
		return nil
	})

	for {
		prompt := promptui.Select{
			Label: "User Management",
			Items: userMenu.GetOptionNames(),
		}

		index, _, err := prompt.Run()
		if err != nil {
			return err
		}

		option := userMenu.GetOption(index)
		if option != nil {
			if option.Name == "Back" {
				return nil
			}

			if err := option.Action(); err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		}
	}
}

func runProjectMenu() error {
	projectMenu := models.NewOptionList()

	// Add project menu options
	projectMenu.AddOption("List Projects", "List all projects", project.ListProjects)
	projectMenu.AddOption("Create Project", "Create a new project", project.CreateProject)
	projectMenu.AddOption("Delete Project", "Delete an existing project", promptDeleteProject)
	projectMenu.AddOption("Back", "Return to main menu", func() error {
		return nil
	})

	for {
		prompt := promptui.Select{
			Label: "Project Management",
			Items: projectMenu.GetOptionNames(),
		}

		index, _, err := prompt.Run()
		if err != nil {
			return err
		}

		option := projectMenu.GetOption(index)
		if option != nil {
			if option.Name == "Back" {
				return nil
			}

			if err := option.Action(); err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		}
	}
}

func promptDeleteUser() error {
	// In a real app, we would fetch actual users from a data source
	users := []string{"User1", "User2", "User3"}

	prompt := promptui.Select{
		Label: "Select user to delete",
		Items: users,
	}

	_, result, err := prompt.Run()
	if err != nil {
		return err
	}

	return user.DeleteUser(result)
}

func promptDeleteProject() error {
	// In a real app, we would fetch actual projects from a data source
	projects := []string{"Project1", "Project2", "Project3"}

	prompt := promptui.Select{
		Label: "Select project to delete",
		Items: projects,
	}

	_, result, err := prompt.Run()
	if err != nil {
		return err
	}

	return project.DeleteProject(result)
}
