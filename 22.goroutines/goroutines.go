package main

import "fmt"

func f(from string) {
	for i := 0; i < 10; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	go f("go routine")

	go func(msg string) {
		for i := 0; i < 10; i++ {
			fmt.Println(msg)
		}
	}("going")

	fmt.Scanln()
	fmt.Println("done")
}
