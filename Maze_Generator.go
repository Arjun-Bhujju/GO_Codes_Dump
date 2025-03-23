package main

import (
    "fmt"
    "math/rand"
    "time"
)

const (
    width  = 10
    height = 10
)

var maze [height][width]bool
var directions = []struct{ x, y int }{{0, -2}, {2, 0}, {0, 2}, {-2, 0}}

func GenerateMaze(x, y int) {
    maze[y][x] = true
    rand.Shuffle(len(directions), func(i, j int) { directions[i], directions[j] = directions[j], directions[i] })
    for _, d := range directions {
        nx, ny := x+d.x, y+d.y
        if ny > 0 && ny < height-1 && nx > 0 && nx < width-1 && !maze[ny][nx] {
            maze[y+d.y/2][x+d.x/2] = true
            GenerateMaze(nx, ny)
        }
    }
}

func PrintMaze() {
    for y := 0; y < height; y++ {
        for x := 0; x < width; x++ {
            if maze[y][x] {
                fmt.Print("  ")
            } else {
                fmt.Print("█ ")
            }
        }
        fmt.Println()
    }
}

func main() {
    rand.Seed(time.Now().UnixNano())
    GenerateMaze(1, 1)
    PrintMaze()
}
