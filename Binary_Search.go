package main

import "fmt"

func BinarySearch(arr []int, target int) int {
    low, high := 0, len(arr)-1
    for low <= high {
        mid := low + (high-low)/2
        if arr[mid] == target {
            return mid
        }
        if arr[mid] < target {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    return -1
}

func main() {
    arr := []int{1, 3, 5, 7, 9, 11, 13}
    fmt.Println("Index of 7:", BinarySearch(arr, 7))
}
