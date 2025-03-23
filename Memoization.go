package main

import "fmt"

var memo = make(map[int]int)

func Fib(n int) int {
    if n <= 1 {
        return n
    }
    if _, found := memo[n]; found {
        return memo[n]
    }
    memo[n] = Fib(n-1) + Fib(n-2)
    return memo[n]
}

func main() {
    fmt.Println("10th Fibonacci number:", Fib(10))
}
