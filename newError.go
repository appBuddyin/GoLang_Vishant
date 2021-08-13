package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

//throw an error if arguments are same
func returnError(a, b int) error {
	if a == b {
		err := errors.New("this is the error creating by errors.new")
		return err
	} else {
		return nil
	}

}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Print("at lest put one argument")
		os.Exit(1)
	}

	arg1, _ := strconv.ParseInt(arguments[1], 10, 64)
	arg2, _ := strconv.ParseInt(arguments[2], 10, 64)
	var err = returnError(int(arg1), int(arg2))
	if err == nil {
		fmt.Println("returnError function run smoothly")
	} else {
		fmt.Println(err)
	}
}
