
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
    drawTree(depth-1, nx, ny, angle-0.5, grid)
    drawTree(depth-1, nx, ny, angle+0.5, grid)
}

func main() {
    size := 20
    grid := make([][]rune, size)
    for i := range grid {
        grid[i] = make([]rune, 40)
        for j := range grid[i] {
            grid[i][j] = ' '
        }
    }

    drawTree(6, 20, 19, -math.Pi/2, grid)

    for _, row := range grid {
        fmt.Println(string(row))
    }
}
