package main

import (
	"fmt"
)

func main() {

	var a = "initial"
	p := fmt.Println
	p(a)

	var b, c int = 1, 2
	p(b, c)

	var d = true
	p(d)

	var e int
	p(e)

	f := "short"
	p(f)
}
