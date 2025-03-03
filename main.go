package main

import (
	"fmt"
	"os"

	actions "github.com/biodigitalJaz/go-cli-terminal/actions"
	controllers "github.com/biodigitalJaz/go-cli-terminal/controllers"
	models "github.com/biodigitalJaz/go-cli-terminal/models"
	"github.com/spf13/cobra"
)

// Define options with associated actions
var options = []models.Option{
	{Text: "Option 1 - Say Hello", Action: actions.PrintHello},
	{Text: "Option 2 - Show Date", Action: actions.ShowDate},
	{Text: "Option 3 - Print Goodbye", Action: actions.PrintGoodby},
	{Text: "Option 4 - Do Nothing", Action: func() {}},
	{Text: "Option 5 - User Input", Action: actions.UserInput},
	{Text: "Exit", Action: func() { os.Exit(0) }},
}

var altOptions = []models.Option{
	{Text: "Alternative 1 - Print Alternative Hello", Action: func() { fmt.Println("Alternative Hello!") }},
	{Text: "Alternative 2 - Show Alternative Date", Action: actions.ShowAlternativeDate},
	{Text: "Alternative 3 - Exit Alternative Menu", Action: func() { os.Exit(0) }},
}

var rootCmd = &cobra.Command{
	Use:   "list-selector",
	Short: "Navigate and select an object from a list using arrow keys",
	Run:   func(cmd *cobra.Command, args []string) { controllers.RunListSelector(options) },
}

var altCmd = &cobra.Command{
	Use:   "alternate",
	Short: "Run the list selector with alternative options",
	Run:   func(cmd *cobra.Command, args []string) { controllers.RunListSelector(altOptions) },
}

func main() {

	rootCmd.AddCommand(altCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
