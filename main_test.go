package main

import (
	"fmt"
	condition "main1/flowControl1/ifElseCondition"
	"testing"
)

//run test in verbose mode
func TestIsEven(t *testing.T) {
	if !condition.IsEven(6) {
		t.Error("Fail 6 is not an even numbers")
	}
}

func TestIntMinBasic(t *testing.T) {
	ans := condition.IntMin(2, -2)
	if ans != -2 {

		t.Errorf("IntMin(2, -2) = %d; expected -2", ans)
	}
}

func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b     int
		expected int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := condition.IntMin(tt.a, tt.b)
			if ans != tt.expected {
				t.Errorf("got %d, expected %d", ans, tt.expected)
			}
		})
	}
}
