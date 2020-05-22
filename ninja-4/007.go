package main

import "fmt"

func main() {
	data := [][]string{{"James", "Bond", "Shaken Not Stirred"}, {"Miss", "MoneyPenny", "Hello, James"}}
	fmt.Println(data)

	for _, i := range data {
		for _, s := range i {
			fmt.Println(s)
		}
	}
}
