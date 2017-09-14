package main 

import "fmt"
import "math"

const s string = "constant"

func main() {
    fmt.Println(s)
    const n = 50002
    const b = n/10
    fmt.Println(b)
    fmt.Println(math.Sin(b))
}