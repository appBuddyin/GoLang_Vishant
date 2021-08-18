package flowControl1

import "fmt"

func Defering() {
	defer fmt.Println("preula")

	fmt.Println("eita")
}
