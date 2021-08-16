package main

import "fmt"

func main() {

	vals := make([]int, 3)

	fmt.Printf("slice: %v; len: %d; cap: %d \n", vals, len(vals), cap(vals))

	vals = append(vals, 1)
	vals = append(vals, 2)
	vals = append(vals, 3)
	vals = append(vals, 4, 5, 6)

	fmt.Printf("slice: %v; len: %d; cap: %d \n", vals, len(vals), cap(vals))
}
