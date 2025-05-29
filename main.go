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

const maxSize = 99999999

type categoryAndMoney struct {
	category string
	money    int
}

var data [maxSize]categoryAndMoney
var count int
var initialBudget int

// Groups
var categoryGroups = map[string]string{
	"Food":      "Food",
	"Snacks":    "Food",
	"Groceries": "Food",
	"Bus":       "Transport",
	"Train":     "Transport",
	"Taxi":      "Transport",
	"Movies":    "Entertainment",
}

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

	editIndexEntry := widget.NewEntry()
	editIndexEntry.SetPlaceHolder("Index to edit")

	editAmountEntry := widget.NewEntry()
	editAmountEntry.SetPlaceHolder("New amount")

	display := widget.NewMultiLineEntry()
	display.Wrapping = fyne.TextWrapWord
	display.SetMinRowsVisible(10)

	groupingInfo := widget.NewLabel(`Category Groupings:
	"Food", "Snacks", "Groceries" → Food
	"Bus", "Train", "Taxi" → Transport
	"Movies" → Entertainment
	Other categories → Other`)

	//Button area
	setBudgetBtn := widget.NewButton("Set Budget", func() {
		setBudgetBtnPressed(budgetEntry, w)
	})


	addExpenseBtn := widget.NewButton("Add Expense", func() {
		addExpenseBtnPressed(categoryEntry, amountEntry, display, w)
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

	sortAlphaBtn := widget.NewButton("Sort A-Z", func() {
		insertionSortByAlphabet(&data, count)
		showData(display)
		dialog.ShowInformation("Sort", "Data sorted alphabetically by category.", w)
	})

	
	editExpenseBtn := widget.NewButton("Edit Amount", func() {
		editExpenseBtnPressed(editIndexEntry, editAmountEntry, display, w)
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
		report += getGroupedReport(&data, count, initialBudget)
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
		container.NewGridWithColumns(2, editIndexEntry, editAmountEntry),
		editExpenseBtn,
		container.NewHBox(selectionSortBtn, insertionSortBtn, sortAlphaBtn, reportBtn),
		groupingInfo,
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
		return 0
	}
	if s[0] == '-' {
		neg = true
		start = 1
	}
	for i := start; i < len(s); i++ {
		ch := s[i]
		if ch < '0' || ch > '9' {
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

func sequentialSearch(data [maxSize]categoryAndMoney, count int, target string) int {
	for i := 0; i < count; i++ {
		if data[i].category == target {
			return i
		}
	}
	return -1
}

func binarySearch(data [maxSize]categoryAndMoney, count int, target int) int {
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

func selectionSort(data *[maxSize]categoryAndMoney, count int) {
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

func insertionSort(data *[maxSize]categoryAndMoney, count int) {
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

// Alphabetical sort
func insertionSortByAlphabet(data *[maxSize]categoryAndMoney, count int) {
	for i := 1; i < count; i++ {
		key := data[i]
		j := i - 1
		for j >= 0 && data[j].category > key.category {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}
}

// Displays the group data
func getGroupedReport(data *[maxSize]categoryAndMoney, count int, budget int) string {
	var groupNames [100]string
	var groupSums [100]int
	groupCount := 0

	for i := 0; i < count; i++ {
		cat := data[i].category
		group, exists := categoryGroups[cat]
		if !exists {
			group = "Other"
		}

		found := false
		idx := 0
		for j := 0; j < groupCount; j++ {
			if groupNames[j] == group {
				found = true
				idx = j
			}
		}

		if found {
			groupSums[idx] += data[i].money
		} else if groupCount < len(groupNames) {
			groupNames[groupCount] = group
			groupSums[groupCount] = data[i].money
			groupCount++
		}
	}

	report := "\nGrouped Category Report:\n"
	for i := 0; i < groupCount; i++ {
		report += fmt.Sprintf("%-15s : %d", groupNames[i], groupSums[i])
		if groupSums[i] > budget/2 {
			report += "  <-- Yeah this is Over 50% of the budget"
		}
		report += "\n"
	}
	return report
}

//Added Functions for the buttons on the main....So its bareable to see for the first few lines :((((. And also to avoid the usage of negative numbers and empty arrays
func addExpenseBtnPressed(categoryEntry *widget.Entry, amountEntry *widget.Entry, display *widget.Entry, w fyne.Window) {
	cat := categoryEntry.Text
	cat = trimSpaces(cat)
	amt := toInt(amountEntry.Text)

	if cat == "" || amt <= 0 {
		dialog.ShowError(fmt.Errorf("cant be less than 0 or 0"), w)
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
}


func editExpenseBtnPressed(editIndexEntry *widget.Entry, editAmountEntry *widget.Entry, display *widget.Entry, w fyne.Window) {
	idx := toInt(editIndexEntry.Text)
	newAmt := toInt(editAmountEntry.Text)
	if idx < 0 || idx >= count {
		dialog.ShowError(fmt.Errorf("invalid index"), w)
		return
	}
	if newAmt <= 0 {
		dialog.ShowError(fmt.Errorf("cant be less than 0 or 0"), w)
		return
	}
	data[idx].money = newAmt
	dialog.ShowInformation("Updated", "Expense updated.", w)
	editIndexEntry.SetText("")
	editAmountEntry.SetText("")
	showData(display)
}


func setBudgetBtnPressed(budgetEntry *widget.Entry, w fyne.Window) {
	val := toInt(budgetEntry.Text)
	if val <= 0 {
		dialog.ShowError(fmt.Errorf("cant be less than 0 or 0"), w)
		return
	}
	initialBudget = val
	dialog.ShowInformation("Budget Set", fmt.Sprintf("Budget set to %d", initialBudget), w)
}


//Manual Space Trimming
func trimSpaces(s string) string {
	start := 0
	for start < len(s) && s[start] == ' ' {
		start++
	}
	end := len(s) - 1
	for end >= 0 && s[end] == ' ' {
		end--
	}
	
	result := ""
	for i := start; i <= end && i < len(s); i++ {
		result += string(s[i])
	}
	return result
}
