package main

import "fmt"

func main() {
	aSlice := []int{1, 2, 3, 4}
	fmt.Println(aSlice)
	integer := make([]int, 2)
	fmt.Println(integer)
	integer = nil
	fmt.Println(integer)

	anArray := [5]int{1, 2, 3, 4, 5}
	refAnArray := anArray[:]
	//changes in anArray will reflect in refAnArray
	fmt.Println(anArray)
	fmt.Println(refAnArray)
	anArray[4] = -11
	fmt.Println(refAnArray)

	s := make([]byte, 5)
	twoD := make([][]int, 4)
	fmt.Println(s)
	fmt.Println(twoD)
	//initialize the elements of twoD slice
	for i := 0; i < len(twoD); i++ {
		for j := 0; j < 4; j++ {
			twoD[i] = append(twoD[i], i*j)
			fmt.Println(twoD[i][j])
		}
	}

	//iterate element though twoD slides
	for _, x := range twoD {
		for _, y := range x {
			print(y, " ")
		}
		fmt.Println()
	}

}
