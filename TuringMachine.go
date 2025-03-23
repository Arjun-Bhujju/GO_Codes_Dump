package main

import "fmt"

type TuringMachine struct {
    tape     []int
    state    int
    position int
}

func (tm *TuringMachine) step() {
    switch tm.state {
    case 0:
        if tm.tape[tm.position] == 0 {
            tm.tape[tm.position] = 1
            tm.position++
            tm.state = 1
        } else {
            tm.position--
            tm.state = 2
        }
    case 1:
        tm.position++
        tm.state = 2
    case 2:
        if tm.tape[tm.position] == 0 {
            tm.state = 3
        } else {
            tm.position--
        }
    case 3:
        return
    }
}

func main() {
    tm := TuringMachine{tape: make([]int, 10)}
    for i := 0; i < 15; i++ {
        tm.step()
        fmt.Println(tm.tape, "State:", tm.state, "Pos:", tm.position)
    }
}
