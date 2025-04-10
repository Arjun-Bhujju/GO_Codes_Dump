package main

import "fmt"

func bubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break // Array is already sorted
		}
	}
	return arr
}

func main() {
	numbers := []int{64, 34, 25, 12, 22, 11, 90}
	sorted := bubbleSort(numbers)
	fmt.Println("Sorted array:", sorted)
}
