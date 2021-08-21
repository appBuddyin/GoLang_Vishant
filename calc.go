package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func cal(w http.ResponseWriter, req *http.Request) {
	expressions, _ := ioutil.ReadAll(req.Body)

	res, err := processStack(string(expressions))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	fmt.Fprintf(w, "the value is %f", res)
}

func main() {
	http.HandleFunc("/", cal)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
func parseArgs(c []string) (float64, float64, error) {
	num1, err := strconv.ParseFloat(c[0], 64)
	if err != nil {
		return 0.0, 0.0, err
	}
	num2, err := strconv.ParseFloat(c[2], 64)
	if err != nil {
		return 0.0, 0.0, err
	}
	return num1, num2, nil
}

func processStack(e string) (float64, error) {
	result := 0.0
	c := strings.Split(e, "")
	if len(c)-1 < 2 {
		return 0.0, errors.New("plz give some args and put space after every token")
	}
	num1, num2, err := parseArgs(c)
	if err != nil {
		return 0.0, err
	}
	switch c[1] {
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0.0 {
			return 0.0, errors.New("can't devide by zero")
		}
		result = num1 / num2
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	}
	return result, nil
}
