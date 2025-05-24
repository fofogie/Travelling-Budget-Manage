package main

import "fmt"

type categoryAndMoney struct {
	category string
	money    int
}

func main() {
	var initialBudget int
	fmt.Print("Enter your starting Budget: ")
	fmt.Scan(&initialBudget)

	var n int
	fmt.Println("Enter how many categories you want to enter first... Note: New categories can be added later on")
	fmt.Scan(&n)

	var data [9999]categoryAndMoney
	var count int

	// Starting input
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

	// Menu
	for {
		fmt.Println("\n--- Menu ---")
		fmt.Println("1. Show all data")
		fmt.Println("2. Add new expense")
		fmt.Println("3. Remove expense by index")
		fmt.Println("4. Search by category (Sequential Search)")
		fmt.Println("5. Search by money (Binary Search)")
		fmt.Println("6. Sort by money (Selection Sort Descending)")
		fmt.Println("7. Sort by money (Insertion Sort Ascending)")
		fmt.Println("8. Show report")
		fmt.Println("0. Exit")
		fmt.Print("Choose: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			showData(&data, count)
		case 2:
			if count >= 9999 {
				fmt.Println("Cannot add more entries. Limit reached.")
				break
			}
			var cath string
			var amount int
			fmt.Println("Enter category:")
			fmt.Scan(&cath)
			fmt.Println("Enter amount:")
			fmt.Scan(&amount)
			data[count] = categoryAndMoney{cath, amount}
			count++
		case 3:
			var idx int
			fmt.Println("Enter index to delete:")
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
		case 4:
			var cat string
			fmt.Println("Enter category to search:")
			fmt.Scan(&cat)
			idx := sequentialSearch(&data, count, cat)
			if idx != -1 {
				fmt.Printf("Found: %s - %d\n", data[idx].category, data[idx].money)
			} else {
				fmt.Println("Not found.")
			}
		case 5:
			var target int
			fmt.Println("Enter money to search:")
			fmt.Scan(&target)
			binarySearch(&data, count, target)
		case 6:
			selectionSort(&data, count)
			fmt.Println("Sorted using Selection Sort (Descending).")
		case 7:
			insertionSort(&data, count)
			fmt.Println("Sorted using Insertion Sort (Ascending).")
		case 8:
			displayData(&data, count, initialBudget)
		case 0:
			return
		default:
			fmt.Println("Invalid Choice")
		}
	}
}

func showData(data *[9999]categoryAndMoney, count int) {
	fmt.Println("\nAll Expenses:")
	for i := 0; i < count; i++ {
		fmt.Printf("Category: %-10s | Amount: %d\n", data[i].category, data[i].money)
	}
}

// Sequential Search jklsjaflksafjklsafjk
func sequentialSearch(data *[9999]categoryAndMoney, count int, target string) int {
	for i := 0; i < count; i++ {
		if data[i].category == target {
			return i
		}
	}
	return -1
}

// Binary Search jahfasdhkjfdshfjkasf
func binarySearch(data *[9999]categoryAndMoney, count int, target int) {
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

// Selection Sort Descending
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

// Insertion Sort Ascending
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

func displayData(data *[9999]categoryAndMoney, count int, plannedMoney int) {
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
