package main

import "fmt"

func GCD(a, b int) int {
    if b == 0 {
        return a
    }
    return GCD(b, a%b)
}

func main() {
    fmt.Println("GCD of 48 and 18:", GCD(48, 18))
}
