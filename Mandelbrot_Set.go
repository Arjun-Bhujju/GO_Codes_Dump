package main

import "fmt"

func Mandelbrot() {
    for y := -12; y <= 12; y++ {
        for x := -30; x <= 30; x++ {
            cx, cy := float64(x)/15.0, float64(y)/8.0
            zx, zy := 0.0, 0.0
            iter := 0
            for zx*zx+zy*zy < 4 && iter < 100 {
                temp := zx*zx - zy*zy + cx
                zy, zx = 2.0*zx*zy+cy, temp
                iter++
            }
            if iter == 100 {
                fmt.Print("█")
            } else {
                fmt.Print(" ")
            }
        }
        fmt.Println()
    }
}

func main() {
    Mandelbrot()
}
