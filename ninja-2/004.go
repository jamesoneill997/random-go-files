package main

import (
	"fmt"
)

var number = 12

func main() {
	fmt.Printf("%d\t%x\t%b\n", number, number, number)

	a := number << 1

	fmt.Printf("%d\t%x\t%b\n", a, a, a)

	fmt.Println("Done")
}
