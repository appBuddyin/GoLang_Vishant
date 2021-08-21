package forLoops

import "fmt"

func ForLoop() {
	sum := 1
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
