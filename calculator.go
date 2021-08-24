package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Operation interface {
	op() float32
}

type variables struct {
	var1 float32
	var2 float32
}

func (v variables) mul() float32 {
	return v.var1 * v.var2
}
func (v variables) add() float32 {
	return v.var1 + v.var2
}
func (v variables) sub() float32 {
	return v.var1 - v.var2
}
func (v variables) dev() float32 {
	return v.var1 / v.var2
}

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
func main() {
	r := mux.NewRouter()

	r.HandleFunc("/calculator/{operation}/{var1:[0-9]+}/{var2:[0-9]+}", ArticleHandler)
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
