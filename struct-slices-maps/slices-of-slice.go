package main

import "fmt"

func main() {
	s1 := make([]int, 5)
	reslice := s1[1:3]
	fmt.Println(reslice)
	fmt.Println(s1)
	for i := 0; i < len(s1); i++ {
		s1[i] = i
	}
	fmt.Println(s1)
	fmt.Println(reslice)
}
