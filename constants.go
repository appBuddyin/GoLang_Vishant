package main

import "fmt"

type Digit int
type power2 int

const PI = 3.14

func main() {
	const s1 = 345
	var v1 float32 = s1 * 4
	fmt.Println(s1)
	fmt.Println(v1)
	fmt.Println(PI)
	const (
		Zero Digit = iota
		One
		Two
		Three
		Four
	)
	fmt.Println(One)
	fmt.Println(Two)
	//using bitwise operator
	const (
		P1 power2 = 1 << iota
		_
		P2
		_
		P3
		_
		P4
	)
	fmt.Println("2^0", P1)
	fmt.Println("2^2", P2)
	fmt.Println("2^4", P3)
	fmt.Println("2^8", P4)
}
