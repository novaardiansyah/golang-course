package main

import "fmt"

func main() {
	printMe("Hello World")

	var numerator int8 = 10
	var denominator int8 = 2

	var result, remainder int8 = intDevision(numerator, denominator)
	fmt.Printf("Result: %v, Remainder: %v\n", result, remainder)
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDevision(numerator int8, denominator int8) (int8, int8) {
	return numerator / denominator, numerator % denominator
}

