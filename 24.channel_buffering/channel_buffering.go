package main

import "fmt"

func main() {
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	go func() {
		fmt.Println(<-messages)
		fmt.Println(<-messages)
	}()
	fmt.Scanln()
}
