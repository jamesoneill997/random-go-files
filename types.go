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

//doesn't change
const helloHello = 11.2

const (
	alpha = iota
	beta  = iota
	delta
)

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
	fmt.Println(helloHello)

	fmt.Println(alpha, beta, delta)

	//bit shifiting
	testBin := 1
	yoda := testBin << 2

	fmt.Println(yoda)

	const (
		_  = iota
		kb = 1 << (iota * 10)
		mb
		gb
	)

	fmt.Println(kb, mb, gb)
}
