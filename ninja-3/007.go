package main

import "fmt"

func main() {
	a := 3
	if a == 2 {
		fmt.Println(a)
	} else if a == 3 {
		fmt.Println(a + 3)
	} else {
		fmt.Println("nope")
	}
}
