package main

import (
	"fmt"
	"os"

	termui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func main() {
	// Initialize termui
	if err := termui.Init(); err != nil {
		panic(err)
	}
	defer termui.Close()

	// Create a paragraph widget to display output
	inputPar := widgets.NewParagraph()
	inputPar.SetRect(0, 0, 50, 3)
	inputPar.Title = "Input"
	termui.Render(inputPar)

	outputPar := widgets.NewParagraph()
	outputPar.SetRect(50, 0, 100, 3)
	outputPar.PaddingTop = 100
	outputPar.Title = "Messages"
	termui.Render(outputPar)

	previous := ""
	current := ""

	// Handle termui events and render the UI
	uiEvents := termui.PollEvents()
    for {
        e := <-uiEvents
		current = current + e.ID
		previous = previous + e.ID
		inputPar.Text = current
		termui.Render(inputPar)
		switch e.ID {
			case "q":
				fmt.Println("q")
				os.Exit(212)
			case  "<c-c":
				fmt.Println("C-c")
				os.Exit(230)
			case "<Enter>": 
				outputPar.Text = previous
				termui.Render(outputPar)
				fmt.Println(previous)
				current = ""
        }
    }
}
