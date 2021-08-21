package flowControl1

import "fmt"

func DeferingMulti() {
	fmt.Println("counting...")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Done!")
}
