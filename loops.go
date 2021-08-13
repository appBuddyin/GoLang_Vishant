package main

import "fmt"

func main() {
	//1st style
	for i := 0; i < 100; i++ {
		if i%30 == 0 {
			continue
		}
		if i == 95 {
			break
		}
		fmt.Print(i, " ")
	}
	fmt.Println()
	//alternate of while loop
	i := 10
	for {
		if i < 0 {
			break
		}
		fmt.Print(i, " ")
		i--
	}
	fmt.Println()
	//alternate of do while loop
	i = 0
	expression := true
	for ok := true; ok; ok = expression {
		if i > 10 {
			expression = false
		}
		fmt.Print(i, " ")
		i++

	}
	fmt.Print()
}
