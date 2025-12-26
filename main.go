package main

import (
	"errors"
	"fmt"
)

func main() {
	printMe("Hello World")

	var numerator int8 = 10
	var denominator int8 = 0

	var result, remainder, err = intDevision(numerator, denominator)
	
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Result: %v, Remainder: %v\n", result, remainder)
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDevision(numerator int8, denominator int8) (int8, int8, error) {
	var err error
	if denominator == 0 {
		err = errors.New("denominator cannot be zero")
		return 0, 0, err
	}
	return numerator / denominator, numerator % denominator, nil
}

