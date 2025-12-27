package main

import (
	"fmt"
)

func main() {
	var myString = []rune("Hello World")
	var indexed = myString[0]
	fmt.Printf("%v, %T\n", indexed, indexed)

	for i, v := range myString {
		fmt.Println(i, v)
	}

	fmt.Printf("\nThe length of 'myString' is %v bytes\n", len(myString))

	var myRune = 'a'
	fmt.Printf("%v, %T\n", myRune, myRune)

	var strSlice = []string{"s", "u", "n", "d", "a", "y"}
	var catStr = ""

	for i := range strSlice {
		catStr += strSlice[i]
	}

	fmt.Printf("%v, %T\n", catStr, catStr)
}