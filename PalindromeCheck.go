package main

import "fmt"

func IsPalindrome(s string) bool {
    n := len(s)
    for i := 0; i < n/2; i++ {
        if s[i] != s[n-1-i] {
            return false
        }
    }
    return true
}

func main() {
    fmt.Println("Is 'racecar' a palindrome?", IsPalindrome("racecar"))
}
