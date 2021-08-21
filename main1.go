package main

import (
	"fmt"
	"main1/flowControl1"
	forLoop "main1/flowControl1/forLoops"
	condition "main1/flowControl1/ifElseCondition"
)

// uncomment the statements and use them
func main() {
	//flow control

	forLoop.ForLoop()

	if condition.IsEven(6) {
		fmt.Print("this is an even number")
	}

	condition.IfElses()

	flowControl1.DeferingMulti()

	flowControl1.Defering()

	flowControl1.Greeting()

	flowControl1.LoopNfunction()

	flowControl1.Switch()

	//	flowControl1.Forever()

}
