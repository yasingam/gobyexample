package main

import "fmt"

func main() {
    
    if 7%2 == 0 { // obviously modulus has a result so
	fmt.Println("7 is even")
    } else {
	fmt.Println("7 is odd") // this will always result with the statement
    }

    if 8%4 == 0 { // again...but yeah, guess its about learning syntax
	fmt.Println("8 is divisible by 4")
    } else {
	fmt.Println("8 is not divisible by 4")
    }

    if 8%2 == 0 || 7%2 == 0 {
	fmt.Println("either 7 or 8 is odd")
    } // no else required, if can stand alone

    if num := 9; num < 0 { // so here assignment statment is before condition
	fmt.Println("num is negative")
    } else if num < 10 { // else if is possible, even if no `else` following
	fmt.Println("num has a single digit")
    } else {
        fmt.Println("num has more than one digit")
    }
	
}
