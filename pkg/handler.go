package pkg

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Calculate(w http.ResponseWriter, r *http.Request) {
	//fetch path
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	operand1, _ := strconv.Atoi(vars["var1"])
	operand2, _ := strconv.Atoi(vars["var2"])
	operands := Variables{
		var1:     float32(operand1),
		var2:     float32(operand2),
		operator: vars["operation"],
	}
	result := operands.op()

	fmt.Fprintf(w, "%f", result)

	// fmt.Fprintf(w, "Category is: %v\n", vars["operation"])
	// fmt.Fprintf(w, "var1 is: %v\n", vars["var1"])
	// fmt.Fprintf(w, "var2 is: %v\n", vars["var2"])
}
