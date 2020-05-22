package main

import "fmt"

func main() {
	states := make([]string, 3, 100)

	states[0] = `Arizona`
	states[1] = `Arkansas`
	states[2] = `California`

	fmt.Println(states)

	for i := 0; i < len(states); i++ {
		fmt.Println(i, states[i])
	}
}
