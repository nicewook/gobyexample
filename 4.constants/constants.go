package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	p := fmt.Println

	p(s)

	const n = 500000000
	const d = 3e20 / n
	p(d)
	p(int64(d))
	p(math.Sin(n))
}
