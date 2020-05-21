package main

import "fmt"

//declare
var number = 3
var word = "hello"
var altWord = `ahh`

//declaring our own type
type hotdog int

var b hotdog

//restrict size of int
var x int64

func main() {

	fmt.Printf("%T\n", number)
	fmt.Println(number)

	//sprintf allows you to print as string, you can assign the value to a variable
	s := fmt.Sprintf(altWord, word, number)
	x = 600
	b = 32

	fmt.Println(s)
	fmt.Printf("%T\n", b)
	fmt.Println(x)

}
