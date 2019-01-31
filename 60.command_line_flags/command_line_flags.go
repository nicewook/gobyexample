package main

import (
	"flag"
	"fmt"
)

func main() {

	wordPtr := flag.String("word", "foo", "a string")

	numPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	p := fmt.Println

	p("word:", *wordPtr)
	p("numb:", *numPtr)
	p("fork:", *boolPtr)
	p("svar:", svar)
	p("tail:", flag.Args())
}
