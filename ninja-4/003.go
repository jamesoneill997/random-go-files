package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y := []int{}
	z := []int{}
	a := []int{}
	b := []int{}

	fmt.Println(x)

	y = append(y, x[0:5]...)
	fmt.Println(y)

	z = append(z, x[5:]...)
	fmt.Println(z)

	a = append(a, x[2:7]...)
	fmt.Println(a)

	b = append(b, x[1:6]...)
	fmt.Println(b)
}
