package main

import (
	"io"
	"os"
)

func main() {
	myString := ""
	var argument []string = os.Args
	if len(argument) == 1 {
		myString = "give at least one argument "
	} else {
		for _, Args := range argument {
			myString = myString + " " + Args
		}
	}

	io.WriteString(os.Stdout, myString)
	io.WriteString(os.Stdout, "\n")

}
