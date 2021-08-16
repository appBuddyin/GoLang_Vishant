package main

import "fmt"

func main() {

	vals := [...]int{1, 2, 3, 4, 5, 6, 7}

	s1 := vals[1:4]
	fmt.Printf("s1: %v, cap: %d\n", s1, cap(s1))

	s2 := vals[5:7]
	fmt.Printf("s2: %v, cap: %d\n", s2, cap(s2))

	s3 := vals[:4]
	fmt.Printf("s3: %v, cap: %d\n", s3, cap(s3))

	s4 := vals[2:]
	fmt.Printf("s4: %v, cap: %d\n", s4, cap(s4))

	s5 := vals[:]
	fmt.Printf("s5: %v, cap: %d\n", s5, cap(s5))
}
