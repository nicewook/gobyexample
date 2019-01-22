package main

import (
	"fmt"
)

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}

}

func intSeq2() func(rst int) int {
	i := 0
	return func(rst int) int {
		i++
		if rst == 1 {
			i = 0
		}
		return i
	}

}

func main() {
	nextInt := intSeq()

	p := fmt.Println

	p(nextInt())
	p(nextInt())
	p(nextInt())

	newInt := intSeq2()

	p(newInt(0))
	p(newInt(0))
	p(newInt(1))
	p(newInt(0))
}
