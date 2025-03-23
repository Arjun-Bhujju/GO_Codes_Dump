package main

import "fmt"

func Sieve(n int) []int {
    primes := make([]bool, n+1)
    for i := 2; i <= n; i++ {
        primes[i] = true
    }
    for p := 2; p*p <= n; p++ {
        if primes[p] {
            for i := p * p; i <= n; i += p {
                primes[i] = false
            }
        }
    }
    result := []int{}
    for i := 2; i <= n; i++ {
        if primes[i] {
            result = append(result, i)
        }
    }
    return result
}

func main() {
    primes := Sieve(50)
    fmt.Println("Prime numbers up to 50:", primes)
}
