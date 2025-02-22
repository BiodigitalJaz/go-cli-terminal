package actions

import (
	"bufio"
	"fmt"
	"os"
)

func UserInput() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your input: ")
	input, _ := reader.ReadString('\n')
	fmt.Printf("You entered: %s", input)
}

func PrintGoodby() {
	fmt.Println("Goodby!")
}

func PrintHello() {
	fmt.Println("Hello, world!")
}

// Function to show the current date
func ShowDate() {
	fmt.Println("Current Date: 2025-02-21") // Replace with actual date logic if needed
}

func ShowAlternativeDate() {
	fmt.Println("Alternative Date: 2025-02-22")
}
