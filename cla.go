package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Println("at least put one argument")
		os.Exit(1)
	}

	arguments := os.Args
	min, _ := strconv.ParseFloat(arguments[1], 64)
	max, _ := strconv.ParseFloat(arguments[1], 64)

	for i := 2; i < len(arguments); i++ {
		floatValue, _ := strconv.ParseFloat(arguments[i], 64)
		if floatValue < min {
			min = floatValue
		}
		if floatValue > max {
			max = floatValue
		}
	}
	fmt.Println("min>", min)
	fmt.Println("max> ", max)
}
