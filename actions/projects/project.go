// actions/project/project.go
package project

import (
	"fmt"
)

// ListProjects lists all projects
func ListProjects() error {
	fmt.Println("Listing all projects...")
	// Implementation would go here
	fmt.Println("Project1\nProject2\nProject3")
	return nil
}

// CreateProject creates a new project
func CreateProject() error {
	fmt.Println("Creating a new project...")
	// Implementation would go here
	fmt.Println("Project created successfully")
	return nil
}

// DeleteProject deletes a project by name
func DeleteProject(name string) error {
	fmt.Printf("Deleting project %s...\n", name)
	// Implementation would go here
	fmt.Println("Project deleted successfully")
	return nil
}
