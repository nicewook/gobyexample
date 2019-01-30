package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	secs := now.Unix()
	secs2 := time.Now().Unix()

	nanos := now.UnixNano()

	fmt.Println(now)
	fmt.Println(secs)
	fmt.Println(secs2)
	fmt.Println()

	millis := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	fmt.Println()
	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}
