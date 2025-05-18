package main

import (
	"fmt"
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Title and start
	a := app.New()
	w := a.NewWindow("Travelling Budget")

	// Title text (centered)
	title := canvas.NewText("Travelling Budget", color.NRGBA{R: 73, G: 149, B: 253, A: 255})
	titleCont := container.New(layout.NewCenterLayout(), title)

	// Text input
	input := widget.NewEntry()
	input.Resize(fyne.NewSize(300, 40))
	inputCont := container.NewMax(input)

	//display area
	displayLabel := widget.NewLabel("")

	//button + button Function
	button := widget.NewButton("Display Input", func() {
		num, err := strconv.Atoi(input.Text)
		if err != nil {
			displayLabel.SetText("Enter Your Expenses")
			return
		}
		result := num + 2
		displayLabel.SetText(fmt.Sprintf("Result: %d", result))
	})

	// Stack title and input vertically
	mainContent := container.NewVBox(
		titleCont,
		inputCont,
		button,
		displayLabel,
	)

	// Window setup
	w.SetContent(mainContent)
	w.Resize(fyne.NewSize(400, 500))
	w.SetFixedSize(true)
	w.ShowAndRun()
}
