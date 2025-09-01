package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	return a / b, nil
}

func add(a, b float64) (float64, error) {
	return a + b, nil
}

func subtract(a, b float64) (float64, error) {
	return a - b, nil
}
func multiply(a, b float64) (float64, error) {
	return a * b, nil
}

func main() {
	var a float64
	fmt.Scanln(&a)

	var b float64
	fmt.Scanln(&b)

	var c string
	fmt.Scanln(&c)

	var res float64
	var err error

	switch c {
	case "+":
		res, err = add(a, b)
	case "-":
		res, err = subtract(a, b)
	case "*":
		res, err = multiply(a, b)
	case "/":
		res, err = divide(a, b)
	}

	if err == nil {
		fmt.Println(res)
	} else {
		fmt.Println(err)
	}

}
