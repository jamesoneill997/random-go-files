package main

import "fmt"

func main() {

	//setting up a slice
	x := []int{2, 3, 4, 5}
	y := []int{51, 52, 53, 45}
	//alternative, sets static size - more efficient if size of array is known
	z := make([]int, 4, 100)
	//multi dimensional
	xm := [][]int{x, y}

	z[0] = 12

	fmt.Println(x)
	fmt.Println(z)
	fmt.Println(len(x))
	fmt.Println(x[0])
	fmt.Println(xm)

	for i, v := range x {
		fmt.Println(i, v)
	}
	//slicing (non inclusive)
	fmt.Println(x[1:3])

	//append to already existing array
	x = append(x, 1, 11, 32)
	fmt.Println(x)
	//merge 2 slices, note the unfurl operator ...
	x = append(x, y...)
	fmt.Println(x)

	//use append to delete elems from slice
	x = append(x[:2], x[4:]...)
	fmt.Println(x)

}
