package controllers

import (
	"fmt"
	"os"

	models "github.com/biodigitalJaz/go-cli-terminal/models"
	"github.com/gdamore/tcell/v2"
)

func RunListSelector(options []models.Option) {
	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating screen: %v\n", err)
		os.Exit(1)
	}

	if err := screen.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing screen: %v\n", err)
		os.Exit(1)
	}
	defer screen.Fini()

	selectedIndex := 0

	// Function to render the list
	render := func() {
		screen.Clear()
		for i, option := range options {
			style := tcell.StyleDefault
			if i == selectedIndex {
				style = style.Background(tcell.ColorBlue).Foreground(tcell.ColorWhite)
			}
			drawText(screen, 2, 2+i, style, option.Text)
		}
		screen.Show()
	}

	render()

	// Event loop for handling keypresses
	for {
		switch ev := screen.PollEvent().(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyUp:
				if selectedIndex > 0 {
					selectedIndex--
				}
			case tcell.KeyDown:
				if selectedIndex < len(options)-1 {
					selectedIndex++
				}
			case tcell.KeyEnter:
				screen.Fini()
				fmt.Printf("Selected: %s\n", options[selectedIndex].Text)
				options[selectedIndex].Action()
				return
			case tcell.KeyEscape, tcell.KeyCtrlC:
				screen.Fini()
				fmt.Println("Exited")
				return
			}
		case *tcell.EventResize:
			screen.Sync()
		}
		render()
	}
}

// drawText renders text at a given position with a style
func drawText(s tcell.Screen, x, y int, style tcell.Style, text string) {
	for i, r := range text {
		s.SetContent(x+i, y, r, nil, style)
	}
}
