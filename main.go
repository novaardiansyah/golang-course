package main

import (
	"fmt"
)

func main() {
	var p *int = new(int)
	var i int = 10

	fmt.Println(p)
	fmt.Println(i)

	fmt.Printf("\nThe value p points to is: %v\n", *p)
	fmt.Printf("The value of i is: %v\n", i)
}
