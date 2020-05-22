package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	for i := range x {
		fmt.Println(x[i])
	}

	fmt.Printf("%T\n", x)
}
