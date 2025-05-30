package main

import "fmt"

const maxSize = 99999999

type categoryAndMoney struct {
	category string
	money    int
}

var categoryGroups = map[string]string{
	"Food":      "Food",
	"Drinks":    "Food",
	"Groceries": "Food",
	"Bus":       "Transport",
	"Train":     "Transport",
	"Taxi":      "Transport",
	"Plane":     "Transport",
	"Parks":     "Entertainment",
	"Movies":    "Entertainment",
}

func main() {
	var initialBudget int
	for initialBudget <= 0 {
		fmt.Println("Enter your starting Budget....Note : Must be more than 0 ")
		fmt.Scan(&initialBudget)
		if initialBudget <= 0 {
			fmt.Println("Budget must be greater than 0.")
		}
	}

	var n int
	for n <= 0 || n > maxSize {
		fmt.Println("Enter how many categories you want to enter first... Note: New categories can be added later on")
		fmt.Scan(&n)
		if n <= 0 || n > maxSize {
			fmt.Println("Enter a number Greater than 0", maxSize)
		}
	}

	var data [maxSize]categoryAndMoney
	var count int
	for i := 0; i < n; i++ {
		var cat string
		var amount int
		fmt.Println("Enter Category:")
		fmt.Scan(&cat)
		fmt.Println("Enter Amount:")
		fmt.Scan(&amount)
		data[count] = categoryAndMoney{cat, amount}
		count++
	}

	exit := false
	for !exit {
		fmt.Println("\n--- Menu ---")
		fmt.Println("1. Show all data")
		fmt.Println("2. Add new expense")
		fmt.Println("3. Remove expense by index")
		fmt.Println("4. Search by category (Sequential Search)")
		fmt.Println("5. Search by money (Binary Search)")
		fmt.Println("6. Sort by money (Selection Sort Descending)")
		fmt.Println("7. Sort by money (Insertion Sort Ascending)")
		fmt.Println("8. Show report")
		fmt.Println("9. Change Expense (Based on Index)")
		fmt.Println("10. Sort by category by alphabet")
		fmt.Println("11. Show grouped category report")
		fmt.Println("0. Exit")
		fmt.Print("Choose: ")

		var choice int
		fmt.Scan(&choice)

		if choice == 1 {
			showData(&data, count)
		} else if choice == 2 {
			if count < 9999 {
				var cath string
				var amount int
				fmt.Println("Enter category:")
				fmt.Scan(&cath)

				for {
					fmt.Println("Enter amount:")
					fmt.Scan(&amount)
					if amount > 0 {
						break
					}
					fmt.Println("Enter a number Greater than 0")
				}
				data[count] = categoryAndMoney{cath, amount}
				count++
			} else {
				fmt.Println("Cannot add more entries. Limit reached.")
			}
		} else if choice == 3 {
			var idx int
			fmt.Print("Enter index to delete: ")
			fmt.Scan(&idx)
			if idx >= 0 && idx < count {
				for i := idx; i < count-1; i++ {
					data[i] = data[i+1]
				}
				count--
				fmt.Println("Expense removed.")
			} else {
				fmt.Println("Invalid index.")
			}
		} else if choice == 4 {
			var target string
			fmt.Print("Enter category to search: ")
			fmt.Scanf(" %[^\n]s", &target)
			idx := sequentialSearch(&data, count, target)
			if idx != -1 {
				fmt.Printf("Found: %s - %d\n", data[idx].category, data[idx].money)
			} else {
				fmt.Println("Not found.")
			}
		} else if choice == 5 {
			var target int
			fmt.Print("Enter money to search: ")
			fmt.Scan(&target)
			binarySearch(&data, count, target)
		} else if choice == 6 {
			selectionSort(&data, count)
			fmt.Println("Sorted using Selection Sort (Descending).")
		} else if choice == 7 {
			insertionSort(&data, count)
			fmt.Println("Sorted using Insertion Sort (Ascending).")
		} else if choice == 8 {
			displayData(&data, count, initialBudget)
		} else if choice == 9 {
			var idx, newAmount int
			fmt.Print("Enter index to edit: ")
			fmt.Scan(&idx)
			if idx >= 0 && idx < count {
				fmt.Printf("Current: %s - %d\n", data[idx].category, data[idx].money)
				for {
					fmt.Print("Enter new amount: ")
					fmt.Scan(&newAmount)
					if newAmount > 0 {
						data[idx].money = newAmount
						fmt.Println("Expense updated.")
						break
					}
					fmt.Println("Enter a number Greater than 0")
				}
			} else {
				fmt.Println("Invalid index.")
			}
		} else if choice == 10 {
			insertionSortByAlphabet(&data, count)
			fmt.Println("Sorted alphabetically.")
		} else if choice == 11 {
			showGroupedReport(&data, count, initialBudget)
		} else if choice == 0 {
			exit = true
		} else {
			fmt.Println("Invalid choice.")
		}
	}
}

func showData(data *[maxSize]categoryAndMoney, count int) {
	fmt.Println("\nAll Expenses:")
	for i := 0; i < count; i++ {
		fmt.Printf("Category: %-10s | Amount: %d\n", data[i].category, data[i].money)
	}
}

func sequentialSearch(data *[maxSize]categoryAndMoney, count int, target string) int {
	for i := 0; i < count; i++ {
		if data[i].category == target {
			return i
		}
	}
	return -1
}

func binarySearch(data *[maxSize]categoryAndMoney, count int, target int) {
	low := 0
	high := count - 1
	for low <= high {
		mid := (low + high) / 2
		if data[mid].money == target {
			fmt.Printf("Found at index %d: %s - %d\n", mid, data[mid].category, data[mid].money)
			return
		} else if data[mid].money < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	fmt.Println("Amount not found")
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

func displayData(data *[maxSize]categoryAndMoney, count int, plannedMoney int) {
	total := 0
	fmt.Println("Report:")
	for i := 0; i < count; i++ {
		fmt.Printf("%s: %d\n", data[i].category, data[i].money)
		total += data[i].money
	}

	fmt.Println("Total:", total)
	diff := plannedMoney - total
	if diff >= 0 {
		fmt.Println("Remaining:", diff)
	} else {
		fmt.Println("Over Budget:", -diff)
	}
}

func showGroupedReport(data *[maxSize]categoryAndMoney, count int, budget int) {
	groupSums := make(map[string]int)
	for i := 0; i < count; i++ {
		cat := data[i].category
		group, exists := categoryGroups[cat]
		if !exists {
			group = "Other"
		}
		groupSums[group] += data[i].money
	}

	fmt.Println("\nGrouped Category Report:")
	for group, sum := range groupSums {
		fmt.Printf("%-15s : %d", group, sum)
		if sum > budget/2 {
			fmt.Printf("  <-- Yeah this is Over 50%% of the budget")
		}
		fmt.Println()
	}
}
