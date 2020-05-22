package main

import "fmt"

func main() {
	//create map between string and int
	m := map[string]int{
		"James": 21,
		"John":  22,
	}

	fmt.Println(m)
	fmt.Println(m["James"])

	//check if k,v pair exists
	if n, a := m["James"]; a {
		fmt.Println(n)
	}

	//add elem to array
	m["johnny"] = 22

	//iterate through map
	for k, v := range m {
		fmt.Println(k, v)
	}
	//remove element from map
	delete(m, "johnny")
	//iterate through map
	for k, v := range m {
		fmt.Println(k, v)
	}

}
