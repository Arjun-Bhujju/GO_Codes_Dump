package main

import "fmt"

func main() {
    source := `package main

import "fmt"

func main() {
    source := %c%s%c
    fmt.Printf(source, 96, source, 96)
}`
    fmt.Printf(source, 96, source, 96)
}
