package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	p := fmt.Println
	pp := fmt.Print
	pp("Write ", i, " as ")
	switch i {
	case 1:
		p("one")
	case 2:
		p("two")
	case 3:
		p("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		p("It's the weekend")
	default:
		p("It's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		p("It's before noon")
	default:
		p("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			p("I'm a bool")
		case int:
			p("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
