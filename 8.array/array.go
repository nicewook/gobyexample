package main

import (
	"fmt"
)

func main() {
	var a [5]int

	p := fmt.Println

	p("emp: ", a)

	a[4] = 100
	p("set: ", a)
	p("get: ", a[4])
	p("len: ", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	p("dcl: ", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	p("2d: ", twoD)
	p("len: ", len(twoD))
}
