package main

import "fmt"

func main() {
	a := 0
	b := 4
	//for loop
	for i := 0; i < 10; i++ {
		fmt.Println("Hello!", i)
	}
	//for statement (while loop equivalent)
	for a <= b {
		fmt.Println("While true")
		if a == 2 {
			fmt.Println("A is 2")
			break
		} else {
			a++
			continue

		}

		switch {
		case false:
			fmt.Println("Don't print this")
			//continue to next

		case true:
			fmt.Println("Print this")
			//go to next case
			fallthrough

		default:
			fmt.Println("default")
		}
	}
}
