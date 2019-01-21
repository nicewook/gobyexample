package main

import "fmt"

func main() {

	p := fmt.Println
	if 7%2 == 0 {
		p("7 is even")
	} else {
		p("7 is odd")
	}

	if 8%4 == 0 {
		p("8 is divisible by 4")
	}

	if num := 9; num < 0 {
		p(num, "is negative")
	} else if num < 10 {
		p(num, "has 1 digit")
	} else {
		p(num, "han multiple digits")
	}
}
