package main

import "fmt"

func main() {

	vals := make([]int, 5, 10)

	n := len(vals)
	c := cap(vals)

	fmt.Printf("The size is: %d\n", n)
	fmt.Printf("The capacity is: %d\n", c)

	vals2 := vals[0:4]
	n2 := len(vals2)
	c2 := cap(vals2)

	fmt.Printf("The size is: %d\n", n2)
	fmt.Printf("The capacity is: %d\n", c2)
}
