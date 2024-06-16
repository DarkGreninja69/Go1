package main

import (
    "fmt"
    "math"
)

const s string = "constant"

func constants() {
    fmt.Println(s)
    const n = 500000000
    fmt.Println(math.Sin(n))
}