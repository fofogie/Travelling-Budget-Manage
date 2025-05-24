package main

import (
	"fmt"

	//all fyne import
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type categoryAndMoney struct {
	category string
	money    int
}

var data [9999]categoryAndMoney
var count int
var initialBudget int

func main() {
	//Display yeah begitulah
	a := app.New()
	w := a.NewWindow("Budget Manager")
	w.Resize(fyne.NewSize(600, 450))

	budgetEntry := widget.NewEntry()
	budgetEntry.SetPlaceHolder("Enter initial budget")

	categoryEntry := widget.NewEntry()
	categoryEntry.SetPlaceHolder("Category name")

	amountEntry := widget.NewEntry()
	amountEntry.SetPlaceHolder("Amount to be spent")

	indexEntry := widget.NewEntry()
	indexEntry.SetPlaceHolder("Index for removal")

	searchCategoryEntry := widget.NewEntry()
	searchCategoryEntry.SetPlaceHolder("Category to search")

	searchAmountEntry := widget.NewEntry()
	searchAmountEntry.SetPlaceHolder("Amount to search")

	display := widget.NewMultiLineEntry()
	display.Wrapping = fyne.TextWrapWord
	display.SetMinRowsVisible(10)

	//buttons area
	setBudgetBtn := widget.NewButton("Set Budget", func() {
		val := toInt(budgetEntry.Text)
		if val == 0 && budgetEntry.Text != "0" {
			dialog.ShowError(fmt.Errorf("invalid budget input"), w)
			return
		}
		initialBudget = val
		dialog.ShowInformation("Budget Set", fmt.Sprintf("Budget set to %d", initialBudget), w)
	})

	addExpenseBtn := widget.NewButton("Add Expense", func() {
		cat := categoryEntry.Text
		amt := toInt(amountEntry.Text)
		if (amt == 0 && amountEntry.Text != "0") || cat == "" {
			dialog.ShowError(fmt.Errorf("invalid category or amount"), w)
			return
		}
		if count >= 9999 {
			dialog.ShowError(fmt.Errorf("data limit reached"), w)
			return
		}
		data[count] = categoryAndMoney{cat, amt}
		count++
		categoryEntry.SetText("")
		amountEntry.SetText("")
		showData(display)
	})

	removeExpenseBtn := widget.NewButton("Remove Expense", func() {
		idx := toInt(indexEntry.Text)
		if idx < 0 || idx >= count {
			dialog.ShowError(fmt.Errorf("invalid index"), w)
			return
		}
		for i := idx; i < count-1; i++ {
			data[i] = data[i+1]
		}
		count--
		indexEntry.SetText("")
		showData(display)
	})

	searchCategoryBtn := widget.NewButton("Search Category", func() {
		cat := searchCategoryEntry.Text
		idx := sequentialSearch(data, count, cat)
		if idx != -1 {
			dialog.ShowInformation("Search Result", fmt.Sprintf("Found: %s - %d", data[idx].category, data[idx].money), w)
		} else {
			dialog.ShowInformation("Search Result", "Category not found.", w)
		}
		searchCategoryEntry.SetText("")
	})

	searchAmountBtn := widget.NewButton("Search Amount", func() {
		target := toInt(searchAmountEntry.Text)
		if target == 0 && searchAmountEntry.Text != "0" {
			dialog.ShowError(fmt.Errorf("invalid amount"), w)
			return
		}
		insertionSort(&data, count)
		idx := binarySearch(data, count, target)
		if idx != -1 {
			dialog.ShowInformation("Search Result", fmt.Sprintf("Found at index %d: %s - %d", idx, data[idx].category, data[idx].money), w)
		} else {
			dialog.ShowInformation("Search Result", "Amount not found.", w)
		}
		searchAmountEntry.SetText("")
	})

	selectionSortBtn := widget.NewButton("Selection Sort", func() {
		selectionSort(&data, count)
		showData(display)
		dialog.ShowInformation("Sort", "Data sorted using Selection Sort (Descending).", w)
	})

	insertionSortBtn := widget.NewButton("Insertion Sort", func() {
		insertionSort(&data, count)
		showData(display)
		dialog.ShowInformation("Sort", "Data sorted using Insertion Sort (Ascending).", w)
	})

	reportBtn := widget.NewButton("Show Report", func() {
		total := 0
		report := "Report:\n"
		for i := 0; i < count; i++ {
			report += fmt.Sprintf("%s: %d\n", data[i].category, data[i].money)
			total += data[i].money
		}
		diff := initialBudget - total
		report += fmt.Sprintf("Total: %d\n", total)
		if diff >= 0 {
			report += fmt.Sprintf("Remaining: %d\n", diff)
		} else {
			report += fmt.Sprintf("Over Budget: %d\n", -diff)
		}
		display.SetText(report)
	})

	//Show app
	w.SetContent(container.NewVBox(
		container.NewGridWithColumns(2, widget.NewLabel("Initial Budget:"), budgetEntry),
		setBudgetBtn,
		container.NewGridWithColumns(2, categoryEntry, amountEntry),
		addExpenseBtn,
		container.NewGridWithColumns(2, indexEntry, removeExpenseBtn),
		container.NewGridWithColumns(2, searchCategoryEntry, searchCategoryBtn),
		container.NewGridWithColumns(2, searchAmountEntry, searchAmountBtn),
		container.NewHBox(selectionSortBtn, insertionSortBtn, reportBtn),
		widget.NewLabel("Data Display:"),
		display,
	))

	w.ShowAndRun()
}

// Manual strconv
func toInt(s string) int {
	n := 0
	neg := false
	start := 0
	if len(s) == 0 {
		fmt.Println("Error: empty input")
		return 0
	}
	if s[0] == '-' {
		neg = true
		start = 1
	}
	for i := start; i < len(s); i++ {
		ch := s[i]
		if ch < '0' || ch > '9' {
			fmt.Println("Error: invalid character")
			return 0
		}
		n = n*10 + int(ch-'0')
	}
	if neg {
		n = -n
	}
	return n
}

func showData(display *widget.Entry) {
	content := "All Entries:\n"
	for i := 0; i < count; i++ {
		content += fmt.Sprintf("%d. %s - %d\n", i, data[i].category, data[i].money)
	}
	display.SetText(content)
}

func sequentialSearch(data [9999]categoryAndMoney, count int, target string) int {
	for i := 0; i < count; i++ {
		if data[i].category == target {
			return i
		}
	}
	return -1
}

func binarySearch(data [9999]categoryAndMoney, count int, target int) int {
	low := 0
	high := count - 1
	for low <= high {
		mid := (low + high) / 2
		if data[mid].money == target {
			return mid
		} else if data[mid].money < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func selectionSort(data *[9999]categoryAndMoney, count int) {
	for i := 0; i < count-1; i++ {
		max := i
		for j := i + 1; j < count; j++ {
			if data[j].money > data[max].money {
				max = j
			}
		}
		if i != max {
			temp := data[i]
			data[i] = data[max]
			data[max] = temp
		}
	}
}

func insertionSort(data *[9999]categoryAndMoney, count int) {
	for i := 1; i < count; i++ {
		key := data[i]
		j := i - 1
		for j >= 0 && data[j].money > key.money {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}
}
