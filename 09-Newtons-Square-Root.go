package main

import (
    "fmt"
    "math"
)

func NewtonsSQRT(x float64) float64 {
    if x == 0 { return 0 }
    z := 1.0
    for i := 0; i < int(x); i++ {
        z = z - ((math.Pow(z, 2) - x) / (2 * z))
    }
    return z
}

func Sqrt(x float64) float64 {
    return math.Sqrt(x)
}

func main() {
    times := 15
    for i := 0; i < times; i++ {
        sqrt := Sqrt(float64(i))
        newt := NewtonsSQRT(float64(i))
        fmt.Println("\n", i, " squared: ")
        fmt.Println("  Square Root: ", sqrt)
        fmt.Println("  Newton's Method: ", newt)
        fmt.Println("  Difference:", math.Abs(sqrt-newt))
    }
}