package main

import (
	"fmt"
)

func main() {
	s := make([]string, 3)
	p := fmt.Println
	p("emp: ", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	p("set: ", s)
	p("get[2]: ", s[2])
	p(len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	p("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	p("cpy: ", c)

	l := s[2:5]
	p("slice[2:5]: ", l)

	l = s[:5]
	p("slice[:5]: ", l)

	l = s[2:]
	p("slice[2:]: ", l)

	t := []string{"g", "h", "i"}

	p("dcl: ", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	p("2d: ", twoD)
}
