package main

import (
    "fmt"
    "math/rand"
    "time"
)

const (
    size = 15
)

var grid [size][size]bool

func Initialize() {
    for i := 0; i < size; i++ {
        for j := 0; j < size; j++ {
            grid[i][j] = rand.Float64() < 0.3
        }
    }
}

func PrintGrid() {
    fmt.Print("\033[H\033[2J") // Clear screen
    for i := 0; i < size; i++ {
        for j := 0; j < size; j++ {
            if grid[i][j] {
                fmt.Print("█ ")
            } else {
                fmt.Print(". ")
            }
        }
        fmt.Println()
    }
    time.Sleep(500 * time.Millisecond)
}

func NextGeneration() {
    var newGrid [size][size]bool
    for i := 0; i < size; i++ {
        for j := 0; j < size; j++ {
            neighbors := 0
            for dx := -1; dx <= 1; dx++ {
                for dy := -1; dy <= 1; dy++ {
                    if dx == 0 && dy == 0 {
                        continue
                    }
                    x, y := i+dx, j+dy
                    if x >= 0 && x < size && y >= 0 && y < size && grid[x][y] {
                        neighbors++
                    }
                }
            }
            newGrid[i][j] = neighbors == 3 || (grid[i][j] && neighbors == 2)
        }
    }
    grid = newGrid
}

func main() {
    rand.Seed(time.Now().UnixNano())
    Initialize()
    for i := 0; i < 50; i++ {
        PrintGrid()
        NextGeneration()
    }
}
