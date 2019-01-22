package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {

	i := 1
	p := fmt.Println
	p("initial:", i)

	zeroval(i)
	p("zeroval:", i)

	zeroptr(&i)
	p("zeroptr:", i)

	p("pointer:", &i)
}
