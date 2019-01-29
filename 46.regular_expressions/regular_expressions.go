package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main(){

	match, _:= regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)
	
	r, _:= regexp.Compile("p([a-z]+)ch")
	fmt.Println(r)
	
	fmt.Println(r.MatchString("peach ddd"))

	fmt.Println(r.FindString("peach punch ddd"))
	fmt.Println(r.FindStringIndex("peach punch ddd"))
	fmt.Println()

	fmt.Println(r.FindStringSubmatch("peach punch ddd"))
	fmt.Println(r.FindStringSubmatchIndex("peach punch ddd"))
	fmt.Println(r.FindAllString("peach punch ddd",-1))
	fmt.Println(r.FindAllStringSubmatch("peach punch ddd",-1))
	fmt.Println(r.FindAllStringSubmatchIndex("peach punch ddd",-1))

	fmt.Println(r.FindAllString("peach punch ddd",2))
	
	fmt.Println(r.Match([]byte("peach")))

  fmt.Println()
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	in:= []byte("a peach")
	out:= r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))

}
