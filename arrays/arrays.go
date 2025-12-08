package main

import "fmt"

func main() {
    var a [5]int // array size declared **before** type
    fmt.Println("emp:", a) // zero-valued

    a[4] = 100 // sets index 4 to 100
    fmt.Println("set:", a)
    fmt.Println("get:", a[4])

    fmt.Println("length:", len(a), "elements")

    b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("declared array:", b)

    b = [...]int{5, 4, 3, 2, 1}
    fmt.Println("reversed:", b)

    c := [...]int{1, 2, 5: 30, 2} // specifies the index of 30
    fmt.Println("zeroes with 30 at position 5:", c)

    // define a 2D array with a for loop
    var twoD [2][3]int // zero-entried 2D array
    for i := range 2 {
	for j := range 3 {
	    twoD[i][j] = i + j
	}
    }
    fmt.Println("for loop-defined 2D array", twoD)

    twoD2 := [2][3]int{{3, 3, 3}, {2, 2, 2}}
    fmt.Println("also 2D:", twoD2)
}
