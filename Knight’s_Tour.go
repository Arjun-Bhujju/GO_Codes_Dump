package main

import (
    "fmt"
)

const N = 8

var moves = [][2]int{
    {2, 1}, {1, 2}, {-1, 2}, {-2, 1},
    {-2, -1}, {-1, -2}, {1, -2}, {2, -1},
}

func isValid(x, y int, board *[N][N]int) bool {
    return x >= 0 && x < N && y >= 0 && y < N && board[x][y] == -1
}

func SolveKT(x, y, move int, board *[N][N]int) bool {
    if move == N*N {
        return true
    }
    for _, m := range moves {
        nx, ny := x+m[0], y+m[1]
        if isValid(nx, ny, board) {
            board[nx][ny] = move
            if SolveKT(nx, ny, move+1, board) {
                return true
            }
            board[nx][ny] = -1
        }
    }
    return false
}

func PrintBoard(board *[N][N]int) {
    for i := 0; i < N; i++ {
        for j := 0; j < N; j++ {
            fmt.Printf("%2d ", board[i][j])
        }
        fmt.Println()
    }
}

func main() {
    board := [N][N]int{}
    for i := range board {
        for j := range board[i] {
            board[i][j] = -1
        }
    }
    board[0][0] = 0
    if SolveKT(0, 0, 1, &board) {
        PrintBoard(&board)
    } else {
        fmt.Println("No solution found")
    }
}
