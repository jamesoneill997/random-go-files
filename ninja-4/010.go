package main

import "fmt"

func main() {
	m := map[string][]string{
		"James": {"Fishing", "Athletics", "Football", "Coding"},
		"Jon":   {"Fishing", "China"},
		"Jack":  {"Football", "Coding"},
	}

	delete(m, "Jon")

	for k, v := range m {
		fmt.Printf("\n%s has interests in:\n", k)
		for i, d := range v {
			fmt.Printf("%d\t%v\n", i+1, d)
		}
	}
}
