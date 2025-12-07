package main

import "fmt"

func main() {
    fmt.Println("go" + "lang") // string concatenation without space
    
    fmt.Println("1 + 1 =", 1 + 1) // adds space after string
    fmt.Println("7.0/3.0 =", 7.0/3.0) // float division
    /* float division requires at least one number to be a float, for example
       7.0/3.0 is equivalent to 7/3.0, 7.0/3, 7/3., 7./3 */
    
    fmt.Println(true&&false) // outputs false since both are not true
    fmt.Println(true||false) // outputs true since one is true
    fmt.Println(!true) // negation of true, therefore false
}
