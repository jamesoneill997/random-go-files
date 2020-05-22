package main

import "fmt"

//create an array, x, containing 5 ints
var x [5]int

func main() {
	//defaults to 0
	fmt.Println(x)
	//set index to a certain value
	x[3] = 100
	fmt.Println(x)

	//get length
	fmt.Println(len(x))

}
