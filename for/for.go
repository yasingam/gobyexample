package main

import "fmt"

func main() {
    i := 1
    for i <=3 { // straightforward conditions
	fmt.Println(i)
	i = i + 1
    }

    for j := 0; j <3; j++ { // C-like, defined variables, conditions, increment
	fmt.Println(j) // prints j then increments at **end** of loop
    }

    for i := range 3 { // i = 0, i = 1, i = 2, exits loop
	fmt.Println(i)
    }

    for i, v := range [5]int{10, 20, 30, 40, 50} { // looping over an array
	    fmt.Println("index:", i, "value:", v) // two vars: index, value
    }

    for { // endless loop
	fmt.Println("loop")
	break // breaks out of endless loop
    }

    for n := range 6 {
	if n%2 == 0 { // Python-like...
	    continue // skips to next iteration of the loop w/out below code
	}
	fmt.Println(n)
    }
}
