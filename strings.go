package main

import "fmt"

var s = "hello"

func main() {
	fmt.Println(s)
	bs := []byte(s)
	fmt.Println(bs)
}
