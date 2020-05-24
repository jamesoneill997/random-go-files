package main

import "fmt"

type person struct{
	first string
	last string
	age int
}

type secretAgent struct{
	person
	ltk bool
}

func main() {
	p1:= person{
		first: "James",
		last: "Bond",
		age:21,
	}

	p2:= secretAgent{
		person{
			first: "James",
			last: "O'Neill",
			age: 21,
		},
		true,
	}

	fmt.Println(p1.age, p2)

}