package main

import "fmt"

type person struct {
	firstName         string
	lastName          string
	favoriteIceCreams [2]string
}

func main() {
	p1 := person{
		firstName:         "John",
		lastName:          "Doe",
		favoriteIceCreams: [2]string{"choc", "mint"},
	}

	p2 := person{
		firstName:         "Paul",
		lastName:          "Bloggs",
		favoriteIceCreams: [2]string{"vanilla", "blueberry"},
	}

	fmt.Println(p1.firstName)

	for i, v := range p1.favoriteIceCreams {
		fmt.Println(i+1, v)
	}
	fmt.Println()

	fmt.Println(p2.firstName)

	for i, v := range p2.favoriteIceCreams {
		fmt.Println(i+1, v)
	}
}
