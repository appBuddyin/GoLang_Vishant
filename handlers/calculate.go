package handlers

import (
	"fmt"
	"net/http"
	"strconv"
)

func calculate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	operand1, _ := strconv.Atoi(vars["var1"])
	operand2, _ := strconv.Atoi(vars["var2"])
	operands := &variables{
		var1: float32(operand1),
		var2: float32(operand2),
	}

	var result float32
	switch vars["operation"] {
	case "add":
		result = operands.add()
	case "mul":
		result = operands.mul()
	case "sub":
		result = operands.sub()
	case "dev":
		result = operands.dev()
	}
	fmt.Fprintf(w, "Category is: %f\n", result)

	// fmt.Fprintf(w, "Category is: %v\n", vars["operation"])
	// fmt.Fprintf(w, "var1 is: %v\n", vars["var1"])
	// fmt.Fprintf(w, "var2 is: %v\n", vars["var2"])
}
