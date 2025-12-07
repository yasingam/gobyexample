package main

import "fmt"

func main() {
    var a = "initial" // assignment without defining type
    fmt.Println(a)

    var b, c int = 1, 2 // assignment with int definition
    fmt.Println(b, c) // adds space between printed variables

    var d = true
    fmt.Println(d) // boolean

    var e int
    fmt.Println(e) // defaults to 'zero-value'

    f := "apple" // shortcut assignment
    fmt.Println(f)
}
