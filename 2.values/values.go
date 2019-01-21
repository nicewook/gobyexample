package main

import (
	"fmt"
)

func main() {
	p := fmt.Println

	p("go" + "lang")
	p("1+1=", 1+1)

	p("7.0/3.0 =", 7.0/3.0)
	p(true && false)
	p(true || false)
	p(!true)
}
