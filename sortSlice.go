package main

import (
	"fmt"
	"sort"
)

type aStruct struct {
	person string
	height int
	weight int
}

func main() {
	mySlice := make([]aStruct, 0)
	mySlice = append(mySlice, aStruct{"shubham", 134, 45})
	mySlice = append(mySlice, aStruct{"hemant", 155, 45})
	mySlice = append(mySlice, aStruct{"naren", 144, 50})
	mySlice = append(mySlice, aStruct{"vishant", 134, 40})
	fmt.Println("0:", mySlice)
	sort.Slice(mySlice, func(i int, j int) bool {
		return mySlice[i].height < mySlice[j].height
	})
	fmt.Print("< in ascending order \n", mySlice, " ")
	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].height > mySlice[j].height
	})
	fmt.Print("\n in descending order> \n", mySlice, " ")
}
