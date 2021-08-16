package main

import "fmt"

func main() {

	words := []string{"falcon", "bold", "bear", "sky", "cloud", "ocean"}

	for idx, word := range words {

		fmt.Println(idx, word)
	}
}
