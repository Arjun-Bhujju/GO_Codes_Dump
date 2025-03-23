package main

import (
    "fmt"
    "time"
)

const size = 11

type Point struct{ x, y int }

func LangtonsAnt(steps int) {
    grid := make(map[Point]bool)
    ant := Point{size / 2, size / 2}
    direction := 0 // 0: up, 1: right, 2: down, 3: left
    moves := []Point{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

    for step := 0; step < steps; step++ {
        if grid[ant] {
            direction = (direction + 3) % 4 // Turn left
        } else {
            direction = (direction + 1) % 4 // Turn right
        }
        grid[ant] = !grid[ant] // Flip the color
        ant.x += moves[direction].x
        ant.y += moves[direction].y

        // Print grid
        fmt.Print("\033[H\033[2J") // Clear screen
        for i := 0; i < size; i++ {
            for j := 0; j < size; j++ {
                if ant.x == i && ant.y == j {
                    fmt.Print("A ")
                } else if grid[Point{i, j}] {
                    fmt.Print("# ")
                } else {
                    fmt.Print(". ")
                }
            }
            fmt.Println()
        }
        time.Sleep(200 * time.Millisecond)
    }
}

func main() {
    LangtonsAnt(50)
}
