// actions/user/user.go
package user

import (
	"fmt"
)

// ListUsers lists all users
func ListUsers() error {
	fmt.Println("Listing all users...")
	// Implementation would go here
	fmt.Println("User1\nUser2\nUser3")
	return nil
}

// CreateUser creates a new user
func CreateUser() error {
	fmt.Println("Creating a new user...")
	// Implementation would go here
	fmt.Println("User created successfully")
	return nil
}

// DeleteUser deletes a user by username
func DeleteUser(username string) error {
	fmt.Printf("Deleting user %s...\n", username)
	// Implementation would go here
	fmt.Println("User deleted successfully")
	return nil
}
