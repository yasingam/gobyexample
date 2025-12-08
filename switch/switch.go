package main

import (
    "fmt"
    "time"
)


func main() {

    i := 4
    fmt.Print(i, " written out is ")
    switch i {
	case 1:
	    fmt.Println("one")
	case 2:
	    fmt.Println("two")
	case 3:
	    fmt.Println("three")
	default:
	    fmt.Println("higher than I can count.")
	}

    t := time.Now()

    switch t.Weekday() {
    case time.Saturday, time.Sunday:
	    fmt.Println("Weekend, baby!")
    default:
	    fmt.Println("It's a week day :(") // no need to escape the (
    }

    switch {
	case t.Hour() < 12:
	    fmt.Println("Good morning!")
	default:
	    fmt.Println("It's after noon!")
    }

    // example with an anonymous function

    whatAmI := func(i interface{}) { // takes in value i of any type
	    switch t := i.(type) { // syntax only legal is swithc statement
	        case int:
		    fmt.Println(t, "is an integer")
		case bool:
		    fmt.Println(t, "is a boolean variable")
		default:
		    fmt.Println("Dunno what", t, "is, so we'll say string")
	    }
    }

    whatAmI(true)
    whatAmI(1)
    whatAmI("hey!")
}
