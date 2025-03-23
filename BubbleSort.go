package main

import "fmt"

func BubbleSort(arr []int) []int {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
    return arr
}

func main() {
    numbers := []int{64, 34, 25, 12, 22, 11, 90}
    fmt.Println("Sorted array:", BubbleSort(numbers))
}
