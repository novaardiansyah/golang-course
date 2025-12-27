package main

import (
	"fmt"
	"strings"
)

func main() {
	var strSlice = []string{"s", "u", "n", "d", "a", "y"}
	var strBuilder strings.Builder

	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	
	var catStr = strBuilder.String()
	fmt.Printf("%v, %T\n", catStr, catStr)
}