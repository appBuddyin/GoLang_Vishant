package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("the compiler is ", runtime.Compiler)
	fmt.Println(runtime.GOARCH, "machine")
	fmt.Println("the compiler is ", runtime.Version())
	fmt.Println("number of cpu ", runtime.NumCPU())
	fmt.Println("number of goroutines ", runtime.NumGoroutine())
	fmt.Println("the operating system is ", runtime.GOOS)
	//get the assembly language by running command
	//$ GOOS=windows GORACH=amd64 go tool compile -S goEnv.go
}
