package main

import (
	"fmt"
)

func main() {
	var myString string = "Hello World"
	var indexed = myString[0]
	fmt.Printf("%v, %T\n", indexed, indexed)

	for i, v := range myString {
		fmt.Println(i, v)
	}

	fmt.Printf("\nThe length of 'myString' is %v bytes\n", len(myString))
}