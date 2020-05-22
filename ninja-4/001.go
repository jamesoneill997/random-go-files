package main

import "fmt"

var arr [5]int

func main() {
	arr[0] = 3
	arr[1] = 4
	arr[2] = 6
	arr[3] = 7
	arr[4] = 8
	for i := range arr {
		fmt.Println(i)
	}

	fmt.Printf("%T\n", arr)
}
