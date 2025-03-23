package main

import (
    "fmt"
    "math"
)

func drawTree(depth, x, y int, angle float64, grid [][]rune) {
    if depth == 0 {
        return
    }
    dx := int(math.Cos(angle) * float64(depth))
    dy := int(math.Sin(angle) * float64(depth))
    nx, ny := x+dx, y-dy
    if nx >= 0 && nx < len(grid[0]) && ny >= 0 && ny < len(grid) {
        grid[ny][nx] = '*'
    }
    drawTree(depth-1, nx, ny, angle-0.4, grid) // Adjusted angle for smoother branches
    drawTree(depth-1, nx, ny, angle+0.4, grid) // Symmetric branching
}

func main() {
    size := 25 // Increased size for a larger tree
    grid := make([][]rune, size)
    for i := range grid {
        grid[i] = make([]rune, 50) // Wider grid for better spacing
        for j := range grid[i] {
            grid[i][j] = ' '
        }
    }

    drawTree(8, 25, 24, -math.Pi/2, grid) // Deeper recursion for more detail

    for _, row := range grid {
        fmt.Println(string(row))
    }
}
