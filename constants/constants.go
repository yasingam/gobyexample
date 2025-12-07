package main

import(
    "fmt"
    "math"
)

const s string = "constant"

func main() {
    fmt.Println(s)

    const n = 50000000000

    const d = 3e20 / n // scientific notation standardd

    fmt.Println(d) // scientific notation

    fmt.Println(int64(d)) // expanded notation

    fmt.Println(math.Sin(n)) // uses radians
}
