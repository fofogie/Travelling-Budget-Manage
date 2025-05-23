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
	fmt.Println("Enter How many category you want to enter first....Note: New categories can be added later on")
	fmt.Scan(&n)

	var data []categoryAndMoney

	// Starting input
	for i := 0; i < n; i++ {
		var cat string
		var amount int
		fmt.Println("Enter Category : ")
		fmt.Scan(&cat)
		fmt.Println("Enter Amount : ")
		fmt.Scan(&amount)
		data = append(data, categoryAndMoney{cat, amount})
	}
	// Menu
	for {
		fmt.Println("\n--- Menu ---")
		fmt.Println("1. Show all data")
		fmt.Println("2. Add new expense")
		fmt.Println("3. Remove expense by index")
		fmt.Println("4. Search by category (Sequential Search)")
		fmt.Println("5. Search by money (Binary Search)")
		fmt.Println("6. Sort by money (Selection Sort)")
		fmt.Println("7. Sort by money (Insertion Sort)")
		fmt.Println("8. Show report")
		fmt.Println("0. Exit")
		fmt.Print("Choose: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			showData(data)
		case 2:
			var cath string
			var amount int
			fmt.Println("Enter category: ")
			fmt.Scan(&cath)
			fmt.Println("Enter amount: ")
			fmt.Scan(&amount)
			data = append(data, categoryAndMoney{cath, amount})
		case 3:
			var idx int
			fmt.Println("Enter index to delete: ")
			fmt.Scan(&idx)

			if idx >= 0 && idx < len(data) {
				for i := idx; i < len(data)-1; i++ {
					data[i] = data[i+1]
				}
				data = data[:len(data)-1]
				fmt.Println("Expense removed.")
			} else {
				fmt.Println("Invalid index.")
			}

		case 4:
			var cat string
			fmt.Println("Enter category to search: ")
			fmt.Scan(&cat)
			idx := sequentialSearch(data, cat)
			if idx != -1 {
				fmt.Printf("Found: %s - %d\n", data[idx].category, data[idx].money)
			} else {
				fmt.Println("Not found.")
			}
		case 5:
			var target int
			fmt.Println("Enter money to search: ")
			fmt.Scan(&target)
			binarySearch(data, target)
		case 6:
			selectionSort(data)
			fmt.Println("Sorted using Selection Sort.")
		case 7:
			insertionSort(data)
			fmt.Println("Sorted using Insertion Sort.")
		case 8:
			displayData(data, initialBudget)
		case 0:
			return
		default:
			fmt.Println("No no no")
		}
	}
}

func showData(data []categoryAndMoney) {
	fmt.Println("\nAll Expenses:")
	for i := 0; i < len(data); i++ {
		fmt.Printf("Category: %-10s | Amount: %d\n", data[i].category, data[i].money)
	}
}

// Search sequent
func sequentialSearch(data []categoryAndMoney, target string) int {
	for i := 0; i < len(data); i++ {
		if data[i].category == target {
			return i
		}
	}
	return -1
}

//search binary
func binarySearch(data []categoryAndMoney, target int) {
	low := 0
	high := len(data) - 1
	for low <= high {
		mid := (low + high) / 2
		if data[mid].money == target {
			fmt.Printf("Found at index %d: %s - %d\n", mid, data[mid].category, data[mid].money)
			return
		} else if data[mid].money < target {
			low = mid + 1
		} else {
			low = mid - 1
		}
	}
	fmt.Print("Ammont not found")

}

// Procedure SelectionSort
func selectionSort(data []categoryAndMoney) {
	n := len(data)
	i := 0
	for i < n-1 {
		min := i
		j := i + 1
		for j < n {
			if data[j].money < data[min].money {
				min = j
			}
			j++
		}
		temp := data[i]
		data[i] = data[min]
		data[min] = temp
		i++
	}
}

//Procedure InsertionSort
func insertionSort(data []categoryAndMoney) {
	i := 1
	for i < len(data) {
		key := data[i]
		j := i - 1
		for j >= 0 && data[j].money > key.money {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
		i++
	}
}

func displayData(data []categoryAndMoney, plannedMoney int) {
	total := 0
	fmt.Println("Report : ")

	for i := 0; i < len(data); i++ {
		fmt.Printf("%s: %d\n", data[i].category, data[i].money)
		total += data[i].money
	}

	fmt.Println("Total: ", total)

	subtrek := plannedMoney - total
	if subtrek >= 0 {
		fmt.Println("Remaining: ", subtrek)
	} else {
		fmt.Println("Over Budget: ", -subtrek)
	}
}
