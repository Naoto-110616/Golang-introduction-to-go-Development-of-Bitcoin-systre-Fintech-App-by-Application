package main

import "fmt"

func main() {

	f := 1.11
	i := int(f)
	fmt.Printf("%T %v\n", i, i)

	// 5,6
	m := map[string]int{
		"Mike":  20,
		"Nancy": 24,
		"Messi": 30,
	}
	fmt.Printf("%T %v", m, m)
}
