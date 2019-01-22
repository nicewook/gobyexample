package main

import (
	"fmt"
)

func main() {

	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	p := fmt.Println

	p("map:", m)

	v1 := m["ki"]
	p("v1:", v1)
	p("len:", len(m))

	delete(m, "k2")
	p("map:", m)

	_, present := m["k2"]
	p("prs2:", present)
	_, present = m["k1"]
	p("prs1:", present)

	n := map[string]int{"foo": 1, "bar": 2}
	p("map:", n)
}
