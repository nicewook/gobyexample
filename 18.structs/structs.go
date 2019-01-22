package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p := fmt.Println

	p(person{"Bob", 20})
	p(person{name: "Alice", age: 20})
	p(person{name: "Freds", age: 20})
	p(&person{name: "Ann", age: 40})

	s := person{name: "Sean", age: 50}
	p(s.name)

	sp := &s
	p(sp.age)

	sp.age = 52
	p(s.age)
	p(sp.age)
}
