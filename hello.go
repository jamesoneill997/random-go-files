package main

import "fmt"

//use var outside function
var outside = 4

//declare int without assigning (defaults to 0 value)
var declare int

func main() {
	//use shorthand in a function
	x := 42
	fmt.Println(x + declare)
	x = 99
	fmt.Println(x + outside)
	hello()
}

func hello() {
	fmt.Println("Hello, World!")
}
