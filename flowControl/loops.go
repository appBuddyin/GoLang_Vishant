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
	fmt.Println()

	anArray := [5]int{0, -2, 1, 3, -1}
	for index, value := range anArray {
		fmt.Println("index: ", index, "value: ", value)
	}
}
