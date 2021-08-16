package main

import (
	"errors"
	"fmt"
)

func main() {
	r, err := divide(100.0, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("result of division is", r)

}

func divide(x, y float32) (float32, error) {
	var result float32

	if y == 0 {
		return result, errors.New("cannot divide by zero")
	}
	result = x / y
	return result, nil 
}