package main

import "fmt"

func main() {
	a := 4

	switch {
	case a == 2:
		fmt.Println("A is 2")
	case a != 2:
		fmt.Println("A is not 2")
	}
}
