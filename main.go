package main

import (
	"strconv"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Calculator")

	// Input fields
	entry1 := widget.NewEntry()
	entry1.SetPlaceHolder("Enter first number")
	entry2 := widget.NewEntry()
	entry2.SetPlaceHolder("Enter second number")

	// Operation dropdown
	operations := []string{"+", "-", "*", "/"}
	operationSelect := widget.NewSelect(operations, func(value string) {})

	// Result label
	result := widget.NewLabel("Result: ")

	// Calculate button
	calculate := widget.NewButton("Calculate", func() {
		num1, err1 := strconv.ParseFloat(entry1.Text, 64)
		num2, err2 := strconv.ParseFloat(entry2.Text, 64)
		op := operationSelect.Selected

		if err1 != nil || err2 != nil {
			result.SetText("Invalid input")
			return
		}

		var res float64
		switch op {
		case "+":
			res = num1 + num2
		case "-":
			res = num1 - num2
		case "*":
			res = num1 * num2
		case "/":
			if num2 != 0 {
				res = num1 / num2
			} else {
				result.SetText("Division by zero")
				return
			}
		default:
			result.SetText("Invalid operation")
			return
		}
		result.SetText("Result: " + strconv.FormatFloat(res, 'f', -1, 64))
	})

	// Layout
	w.SetContent(container.NewVBox(
		entry1,
		entry2,
		operationSelect,
		calculate,
		result,
	))

	w.ShowAndRun()
}
