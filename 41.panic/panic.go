package main

import "os"

func main() {
	panic("a problem")

	os.Create("./panic.txt")
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
