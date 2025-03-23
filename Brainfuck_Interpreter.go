package main

import (
    "bufio"
    "fmt"
    "os"
)

func interpretBrainfuck(code string) {
    memory := make([]byte, 30000)
    ptr := 0
    loopStack := []int{}
    
    for i := 0; i < len(code); i++ {
        switch code[i] {
        case '>':
            ptr++
        case '<':
            ptr--
        case '+':
            memory[ptr]++
        case '-':
            memory[ptr]--
        case '.':
            fmt.Printf("%c", memory[ptr])
        case ',':
            reader := bufio.NewReader(os.Stdin)
            input, _ := reader.ReadByte()
            memory[ptr] = input
        case '[':
            if memory[ptr] == 0 {
                depth := 1
                for depth > 0 {
                    i++
                    if code[i] == '[' {
                        depth++
                    } else if code[i] == ']' {
                        depth--
                    }
                }
            } else {
                loopStack = append(loopStack, i)
            }
        case ']':
            if memory[ptr] != 0 {
                i = loopStack[len(loopStack)-1] - 1
            } else {
                loopStack = loopStack[:len(loopStack)-1]
            }
        }
    }
}

func main() {
    brainfuckCode := "++++++++++[>+>+++>+++++++>++++++++++<<<<-]>>>++.>+.+++++++..+++."
    interpretBrainfuck(brainfuckCode)
}
